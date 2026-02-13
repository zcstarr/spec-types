import * as path from "path";
import Transpiler from "@json-schema-tools/transpiler";
import { compileTypescript, StringUtils } from "./util";
import * as fs from "fs";
import { promisify } from "util";
import type { JSONSchemaObject } from "@json-schema-tools/meta-schema";
import Dereferencer from "@json-schema-tools/dereferencer";
import toml from "@iarna/toml";
import {getAllSchemas} from "test-open-rpc-spec"

const readFile = promisify(fs.readFile);
const writeFile = promisify(fs.writeFile);
const mkdir = promisify(fs.mkdir);

import ts from "typescript";
import { config } from "process";

let transpilerCache: Record<string, Transpiler> = {};
interface GetTranspiler {
  (name: string): Transpiler
}

const buildTranspilerCache = async (schemas: Record<string, any>): Promise<GetTranspiler> => {
  const cache: Record<string, Transpiler> = {};
  for(const [name, schema] of Object.entries(schemas)) {
    try {
      const dereffer = new Dereferencer(JSON.parse(JSON.stringify(schema)));
      const dereffedSchema = await dereffer.resolve();
      cache[name] = new Transpiler(dereffedSchema);
    }catch(e: any) {
      throw new Error(`Failed to get transpiler for ${name}: ${e.message}`);
    }
  }
  return (name: string) => {
    const transpiler = cache[name];
    if(!transpiler) {
      throw new Error(`Transpiler for ${name} not found`);
    }
    return transpiler;
  }
}

type Op = 
  | {
    "type": "write"; path: string; content: string;
  }
  | {
    "type": "mkdir"; path: string;
  }
  | {
    "type": "compile"; fileNames: string[]; options: ts.CompilerOptions | any, lang: "ts" | "go" | "rs" | "py";
  }



const generateTsOp = (getTranspiler: GetTranspiler, schemasNames: string[] , outpath: string): Op[] => {
  const ops: Op[] = [{ type: "mkdir", path: outpath }]

  return ops.concat(schemasNames.flatMap((name) => {
    return [
      { type: "mkdir", path: `${outpath}/${name}` },
      { type: "write", path: `${outpath}/${name}/index.ts`, content: getTranspiler(name).toTs() }
    ];
  })).concat([
    { type: "compile", fileNames: [outpath], options: {
      target: ts.ScriptTarget.ES2022,
      module: ts.ModuleKind.ES2022,
      lib: ["lib.es2022.d.ts"],
      moduleResolution: ts.ModuleResolutionKind.Bundler,
      declaration: true,
      outDir: outpath,
      strict: true,
    }, lang: "ts" }
  ]);
}

// Interpreter does the actual work
const execute = async (ops: Op[]) => {
  for (const op of ops) {
    switch (op.type) {
      case "write": await writeFile(op.path, op.content); break;
      case "mkdir": await mkdir(op.path, { recursive: true }); break;
      case "compile":  {
        switch (op.lang) {
          case "ts": compileTypescript(op.fileNames, op.options as ts.CompilerOptions); break;
          default: throw new Error(`Unsupported language: ${op.lang}`);
        }
      }
    }
  }
};

const run = async () => {

  const getTranspiler = await buildTranspilerCache(getAllSchemas());
  const ops = generateTsOp(getTranspiler, Object.keys(getAllSchemas()), "./generated/packages/ts");
  await execute(ops);
}

run()

const generateTs = async (transpiler: Transpiler, schema: JSONSchemaObject, outpath: string): Promise<boolean> => {
  const indexTS = `${outpath}/src/index.ts`;
  const regularName = StringUtils.camelCase(schema.title);
  const tsCode = [
    `export const ${regularName} = ${JSON.stringify(schema)};`,
    `export default ${regularName}`
  ].join("\n");
  await writeFile(indexTS, tsCode);
  const program = ts.createProgram([indexTS, `${outpath}/src/schema.json`], {
    target: ts.ScriptTarget.ES2022,
    module: ts.ModuleKind.ES2022,
    lib: ["lib.es2022.d.ts"],
    moduleResolution: ts.ModuleResolutionKind.Bundler,
    declaration: true,
    outDir: outpath,
    strict: true,
    esModuleInterop: true,
    resolveJsonModule: true,
  });
  const emitResult = program.emit();
  const diagnostics = ts.getPreEmitDiagnostics(program).concat(emitResult.diagnostics);
  if (diagnostics.length > 0) {
    const msg = ts.formatDiagnosticsWithColorAndContext(diagnostics, {
      getCanonicalFileName: (f) => f,
      getCurrentDirectory: ts.sys.getCurrentDirectory,
      getNewLine: () => ts.sys.newLine,
    });
    throw new Error(msg);
  }
  await writeFile(`${outpath}/index.d.ts`, transpiler.toTs());
  return true;
}

const generateGo = async (transpiler: Transpiler, schema: JSONSchemaObject, outpath: string): Promise<boolean> => {
  const packageName = StringUtils.snakeCase(schema.title);
  const exportName = `Raw${StringUtils.upperFirst(packageName)}`;
  const escapedS = JSON.stringify(schema).replace(/"/g, "\\\"");

  const go = [
    `package ${packageName}`,
    "",
    "",
    transpiler.toGo(),
    "",
    `const ${exportName} = "${escapedS}"`,
  ].join("\n");

  await writeFile(`${outpath}/${packageName}.go`, go);
  return true;
}

const generateRs = async (transpiler: Transpiler, schema: JSONSchemaObject, outpath: string, version: string): Promise<boolean> => {
  const crateName = StringUtils.snakeCase(schema.title);

  let cargotoml;

  try {
    cargotoml = toml.parse(await readFile("${outpath}/Cargo.toml", "utf8"));
    cargotoml.version = version;
  } catch (e) {
    cargotoml = {
      package: {
        name: crateName,
        version,
        description: "Generated types based on the JSON-Schema for " + crateName,
        license: "Apache-2.0"
      },
      dependencies: {
        serde: { version: "1.0", features: ["derive"] },
        serde_json: "1.0", // eslint-disable-line
        derive_builder: "0.10" // eslint-disable-line
      }
    }
  }

  await writeFile(`${outpath}/Cargo.toml`, toml.stringify(cargotoml));
  await writeFile(`${outpath}/src/lib.rs`, transpiler.toRs());

  return true;
};

export const prepare = async (): Promise<boolean> => {

  const outpath = pluginConfig.outpath || process.cwd();
  await mkdir(`./${outpath}/src`, { recursive: true });

  const schemaPath = path.resolve(process.cwd(), pluginConfig.schemaLocation);

  const schemaString = await readFile(schemaPath, "utf8");
  const schema = JSON.parse(schemaString);

  await writeFile(`${outpath}/src/schema.json`, schemaString);

  if (!schema.title) {
    throw new Error("The schema must have a title", "ENOTITLE", "Schema requires a title");
  }

  let dereffedSchema;
  try {
    const dereffer = new Dereferencer(JSON.parse(JSON.stringify(schema)));
    dereffedSchema = await dereffer.resolve();
  } catch (e: any) {
    throw new Error(e.message);
  }

  const transpiler = new Transpiler(dereffedSchema);

  if (config.languages?.ts) {
    await generateTs(transpiler, schema, outpath);
  }
  if (config.languages?.go) {
    await generateGo(transpiler, schema, outpath);
  }
  if (config.languages?.rs) {
    await generateRs(transpiler, schema, outpath, context.nextRelease.version)
  }
  if (config.languages?.py) {
    await writeFile(`${outpath}/index.py`, transpiler.toPy());
  }

  return true;
}
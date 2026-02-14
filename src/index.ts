import * as path from "path";
import Transpiler from "@json-schema-tools/transpiler";
import { compileTypescript, StringUtils } from "./util";
import { buildPackageJson, buildCargoToml, buildGoMod, buildPyProjectToml } from "./assets.ts";
import {readFile, writeFile, mkdir, rmdir, rm} from "fs/promises";
import type { JSONSchemaObject } from "@json-schema-tools/meta-schema";
import Dereferencer from "@json-schema-tools/dereferencer";
import toml from "@iarna/toml";
import {getAllSchemas} from "test-open-rpc-spec"
import ts from "typescript";


interface GetTranspiler {
  (name: string): Transpiler
}

// Operation types for the system 
type Op = 
  | {
    "type": "write"; path: string; content: string;
  }
  | {
    "type": "rm"; path: string;
  }
  | {
    "type": "mkdir"; path: string;
  }
  | {
    "type": "compile"; fileNames: string[]; options: ts.CompilerOptions | any, lang: "ts" | "go" | "rs" | "py";
  }


// Programatically construct all the assets tomls pyprojects  generate them 
// then we will use Knope to have the changesets and versions propogated and committed 
// this will then allow us to have nice generated assets that can be used in the spec like changesets

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

// Index file generators

const tsIndexFile = (schemaNames: string[]): string => {
  const imports = schemaNames
    .map((name) => `import type * as V${name} from "./${name}/index.js";`);

  const reexport = `export type { ${schemaNames.map((name) => `V${name}`).join(", ")} }`;

  return [...imports, "", reexport].join("\n");
};

const rsLibFile = (schemaNames: string[]): string => {
  return schemaNames.map((name) => `pub mod v${name};`).join("\n") + "\n";
}

const goPackageFile = (name: string, goCode: string): string => {
  return `package v${name}\n\n${goCode}\n`;
}

const pyInitFile = (schemaNames: string[]): string => {
  return schemaNames.map((name) => `from . import v${name}`).join("\n") + "\n";
}

// Operations generators

const generateTsOp = (getTranspiler: GetTranspiler, schemasNames: string[] , outpath: string): Op[] => {
  const ops: Op[] = [{ type: "rm", path: `${outpath}` }, { type: "mkdir", path: outpath }]
  return ops.concat(schemasNames.flatMap((name) => {
    return [
      { type: "mkdir", path: `${outpath}/${name}` },
      { type: "write", path: `${outpath}/${name}/index.ts`, content: getTranspiler(name).toTs() },
      { type: "write", path: `${outpath}/index.ts`, content: tsIndexFile(schemasNames, outpath) }
    ];
  })).concat([
    { type: "write", path: `${outpath}/package.json`, content: JSON.stringify(buildPackageJson(schemasNames), null, 2) },
    { type: "compile", fileNames: [
      `${outpath}/index.ts`,
      ...schemasNames.map((name) => `${outpath}/${name}/index.ts`),
    ], options: {
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

const generateRsOp = (getTranspiler: GetTranspiler, schemasNames: string[], outpath: string): Op[] => {
  const ops: Op[] = [
    { type: "rm", path: outpath },
    { type: "mkdir", path: outpath },
    { type: "mkdir", path: `${outpath}/src` },
  ];

  return ops
    .concat(
      schemasNames.map((name) => ({
        type: "write" as const,
        path: `${outpath}/src/v${name}.rs`,
        content: getTranspiler(name).toRs(),
      })),
    )
    .concat([
      {
        type: "write",
        path: `${outpath}/src/lib.rs`,
        content: rsLibFile(schemasNames),
      },
      {
        type: "write",
        path: `${outpath}/Cargo.toml`,
        content: toml.stringify(buildCargoToml(schemasNames) as any),
      },
    ]);
}

const generateGoOp = (getTranspiler: GetTranspiler, schemasNames: string[], outpath: string): Op[] => {
  const ops: Op[] = [
    { type: "rm", path: outpath },
    { type: "mkdir", path: outpath },
  ];

  return ops
    .concat(
      schemasNames.flatMap((name) => [
        { type: "mkdir" as const, path: `${outpath}/v${name}` },
        {
          type: "write" as const,
          path: `${outpath}/v${name}/v${name}.go`,
          content: goPackageFile(name, getTranspiler(name).toGo()),
        },
      ]),
    )
    .concat([
      {
        type: "write",
        path: `${outpath}/go.mod`,
        content: buildGoMod(),
      },
    ]);
}

const generatePyOp = (getTranspiler: GetTranspiler, schemasNames: string[], outpath: string): Op[] => {
  const pkg = "open_rpc_spec_types";
  const ops: Op[] = [
    { type: "rm", path: outpath },
    { type: "mkdir", path: outpath },
    { type: "mkdir", path: `${outpath}/src/${pkg}` },
  ];

  return ops
    .concat(
      schemasNames.map((name) => ({
        type: "write" as const,
        path: `${outpath}/src/${pkg}/v${name}.py`,
        content: getTranspiler(name).toPy(),
      })),
    )
    .concat([
      {
        type: "write",
        path: `${outpath}/src/${pkg}/__init__.py`,
        content: pyInitFile(schemasNames),
      },
      {
        type: "write",
        path: `${outpath}/pyproject.toml`,
        content: toml.stringify(buildPyProjectToml() as any),
      },
    ]);
}


// Interpreter does the actual work
const execute = async (ops: Op[]) => {
  for (const op of ops) {
    switch (op.type) {
      case "write": await writeFile(op.path, op.content); break;
      case "mkdir": await mkdir(op.path, { recursive: true }); break;
      case "rm": await rm(op.path, { recursive: true, force: true }); break;
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
  const schemaNames = Object.keys(getAllSchemas());

  const tsOps = generateTsOp(getTranspiler, schemaNames, "./generated/packages/ts");
  const rsOps = generateRsOp(getTranspiler, schemaNames, "./generated/packages/rs");
  const goOps = generateGoOp(getTranspiler, schemaNames, "./generated/packages/go");
  const pyOps = generatePyOp(getTranspiler, schemaNames, "./generated/packages/py");

  await Promise.all([
    execute(tsOps),
    execute(rsOps),
    execute(goOps),
    execute(pyOps),
  ]);
}

run()
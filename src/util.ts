import * as ts from "typescript";

export const StringUtils = {
  upperFirst: (str?: string): string =>
    str ? str.charAt(0).toUpperCase() + str.slice(1) : "",

  camelCase: (str?: string): string =>
    str
      ? str
          .toLowerCase()
          .replace(/[^a-zA-Z0-9]+(.)/g, (m, chr) => chr.toUpperCase())
      : "",

  snakeCase: (str?: string): string =>
    str
      ? str
          .match(
            /[A-Z]{2,}(?=[A-Z][a-z]+[0-9]*|\b)|[A-Z]?[a-z]+[0-9]*|[A-Z]|[0-9]+/g,
          )!
          .map((x: string) => x.toLowerCase())
          .join("_")
      : "",
};

export const compileTypescript = (
  fileNames: string[],
  options: ts.CompilerOptions,
): void => {
  const program = ts.createProgram(fileNames, options);
  const emitResult = program.emit();

  const allDiagnostics = ts
    .getPreEmitDiagnostics(program)
    .concat(emitResult.diagnostics);

  allDiagnostics.forEach((diagnostic) => {
    const message = ts.flattenDiagnosticMessageText(
      diagnostic.messageText,
      "\n",
    );
    // eslint-disable-next-line no-console
    console.log(`${diagnostic.file?.fileName}: ${message}`);
  });
};

// Usage
/*
compile(['src/index.ts'], {
  noEmitOnError: true,
  target: ts.ScriptTarget.ESNext,
  module: ts.ModuleKind.CommonJS,
  outDir: './dist'
});
*/

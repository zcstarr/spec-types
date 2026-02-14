export interface PackageJsonOptions {
  name: string;
  version: string;
}

export interface CargoTomlOptions {
  name: string;
  version: string;
}

export const buildCargoToml = (
  schemaNames: string[],
  opts: CargoTomlOptions = { name: "open-rpc-spec-types", version: "0.0.0" },
): object => ({
  package: {
    name: opts.name,
    version: opts.version,
    edition: "2021",
    description: "Generated OpenRPC specification types",
    license: "Apache-2.0",
  },
  dependencies: {
    serde: { version: "1.0", features: ["derive"] },
    serde_json: "1.0",
    derive_builder: "0.20",
  },
});

export interface GoModOptions {
  module: string;
  goVersion: string;
}

export const buildGoMod = (
  opts: GoModOptions = { module: "github.com/open-rpc/spec-types", goVersion: "1.24.5" },
): string => `module ${opts.module}\n\ngo ${opts.goVersion}\n`;

export interface PyProjectTomlOptions {
  name: string;
  version: string;
}

export const buildPyProjectToml = (
  opts: PyProjectTomlOptions = { name: "open-rpc-spec-types", version: "0.0.0" },
): object => ({
  "build-system": {
    requires: ["hatchling"],
    "build-backend": "hatchling.build",
  },
  project: {
    name: opts.name,
    version: opts.version,
    description: "Generated OpenRPC specification types",
    license: "Apache-2.0",
    "requires-python": ">=3.12",
  },
});

export const buildPackageJson = (
  schemaNames: string[],
  opts: PackageJsonOptions = { name: "@open-rpc/spec-types", version: "0.0.0" },
) => {
  const subpathExports = Object.fromEntries(
    schemaNames.map((name) => [
      `./${name}`,
      {
        bun: `./${name}/index.ts`,
        types: `./${name}/index.d.ts`,
        import: `./${name}/index.js`,
        default: `./${name}/index.js`,
      },
    ]),
  );

  return {
    name: opts.name,
    version: opts.version,
    type: "module",
    module: "index.ts",
    exports: {
      ".": {
        bun: "./index.ts",
        types: "./index.d.ts",
        import: "./index.js",
        default: "./index.js",
      },
      ...subpathExports,
    },
    files: ["**/*.ts", "**/*.d.ts", "**/*.js", "!node_modules"],
  };
};

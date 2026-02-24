export interface PackageJsonOptions {
  name: string;
  version: string;
  dependencies: Record<string, string>;
}

export interface CargoTomlOptions {
  name: string;
  version: string;
}

export const buildCargoToml = (
  schemaNames: string[],
  opts: CargoTomlOptions = { name: "@open-rpc/spec-types", version: "0.0.0" },
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
  opts: GoModOptions = { module: "github.com/zcstarr/spec-types/generated/packages/go", goVersion: "1.24.5" },
  version?: string,
): string => {
  const versionComment = version ? ` // ${version}` : "";
  return `module ${opts.module}${versionComment}\n\ngo ${opts.goVersion}\n`;
};

export interface PyProjectTomlOptions {
  name: string;
  version: string;
}

export const buildPyProjectToml = (
  opts: PyProjectTomlOptions = { name: "@open-rpc/spec-types", version: "0.0.0" },
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

export const buildTsConfig = () => ({
  compilerOptions: {
    target: "ESNext",
    module: "NodeNext",
    moduleResolution: "NodeNext",
    declaration: true,
    emitDeclarationOnly: false,
    strict: true,
    skipLibCheck: true,
  },
  include: ["./**/*.ts"],
});

export const buildPackageJson = (
  schemaNames: string[],
  opts: PackageJsonOptions = { name: "@open-rpc/spec-types", version: "0.0.0", dependencies: {} },
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
    scripts: {
      build: "./node_modules/.bin/tsc",
      prepack: "./node_modules/.bin/tsc",
    },
    exports: {
      ".": {
        bun: "./index.ts",
        types: "./index.d.ts",
        import: "./index.js",
        default: "./index.js",
      },
      ...subpathExports,
    },
    dependencies: opts.dependencies,
    devDependencies: {
      typescript: "^5.0.0",
    },
    files: ["**/*.ts", "**/*.d.ts", "**/*.js", "!node_modules"],
  };
};

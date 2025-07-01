# ksl-schema-language

The KSL (Kessel Schema Language) is a domain-specific language for defining authorization schemas.

## Installation

### From Source (Go Install)

You can install the `ksl` command-line tool directly from the GitHub repository:

```bash
go install github.com/project-kessel/ksl-schema-language/cmd/ksl@latest
```

This will install the `ksl` binary to your `$GOPATH/bin` directory (or `$HOME/go/bin` if `GOPATH` is not set).

### From Source (Manual Build)

Alternatively, you can clone the repository and build it locally:

```bash
git clone https://github.com/project-kessel/ksl-schema-language.git
cd ksl-schema-language
make build
```

This will create the binary in the `./bin/` directory.

## Usage

The `ksl` tool can be used to:

1. **Compile KSL files to intermediate representation:**
   ```bash
   ksl -c input.ksl -o output.json
   ```

2. **Build multiple KSL/JSON files into a SpiceDB schema:**
   ```bash
   ksl file1.ksl file2.json -o schema.zed
   ```

### Command-line Options

- `-c <file>`: Compile a KSL source file to intermediate representation
- `-o <file>`: Specify the output file (defaults to `schema.zed` for builds or `<input>.json` for compilation)

## Examples

See the `samples/` and `demo/` directories for example KSL files and usage patterns.
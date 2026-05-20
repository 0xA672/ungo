# ungo

Safely uninstall any Go toolchain installation with automatic detection and interactive confirmation.

## Why ungo?

Most Go installation guides tell you how to install, but not how to remove. `ungo` fills that gap. It locates your Go installation (no matter how it was installed), asks for confirmation, removes the installation directory, and tells you how to clean up your `PATH`.

## Features

- **Automatic detection** – finds Go via `$GOROOT`, `go env GOROOT`, or common default paths
- **Safe and interactive** – asks for explicit confirmation before deleting anything
- **Cross‑platform** – works on Linux, macOS, and Windows
- **No dependencies** – single Go binary, no runtime required

## Installation

### Build from source (recommended)

You need Go installed to build `ungo` (yes, it's a bit ironic – use it to uninstall a different Go installation, or build it on another machine).

```bash
git clone https://github.com/0xA672/ungo.git
cd ungo
go build -o ungo main.go
```

Then move the `ungo` binary to a location in your `PATH` if you like.

### Pre-built binaries

*(You can publish binaries via GitHub Releases when ready.)*

## Usage

```bash
./ungo
```

`ungo` will:

1. Detect your Go installation location
2. Verify that it really looks like a Go installation
3. Ask you to type `yes` to confirm deletion
4. Remove the entire installation directory
5. Print clear instructions for removing the Go binary directory from your `PATH`

### Example

```
$ ./ungo
=== ungo: Go toolchain uninstaller ===
Found Go installation at: /usr/local/go
Are you sure you want to delete this Go installation? This cannot be undone! (yes/no): yes
Removing /usr/local/go ...
Go installation directory removed successfully.

Warning: you must manually remove the following directory from your PATH environment variable:
  /usr/local/go/bin
...
Uninstall complete. Go has been removed from your system.
```

## Important notes

- `ungo` **does not** modify your shell configuration files (`.bashrc`, `.zshrc`, etc.) automatically. It tells you exactly which line to remove – this is safer and avoids breaking your setup.
- If Go is installed in a system‑protected location (e.g., via a package manager or Nix), `ungo` will fail to delete it with a permission error. Use the appropriate package manager to uninstall in those cases.

## License

[MIT](LICENSE)

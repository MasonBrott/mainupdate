# mainupdate

A simple Go tool to update your local default branch from origin.

## What it does

1. Checks that your working tree is clean
2. Fetches all branches and tags from origin and prunes any local branches that no longer exist on the remote
3. Detects the repository's default base branch (e.g. `main`, `master`, or another)
4. Checks out the default branch
5. Pulls the latest changes from origin

## Usage

```bash
go run main.go
```

Or build and run:

```bash
go build -o mainupdate
./mainupdate
```

## Building small executables

Cross-compile from any machine with Go installed. Use `-ldflags="-s -w"` to strip debug symbols and shrink the binary.

```bash
mkdir -p dist

# Linux (amd64)
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o dist/mainupdate-linux-amd64 .

# Windows (amd64)
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o dist/mainupdate-windows-amd64.exe .

# macOS (Apple Silicon)
GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o dist/mainupdate-darwin-arm64 .
```

Copy the file for your platform onto your `PATH` (or run it directly). On macOS, unsigned binaries from cross-compilation may need to be allowed under **System Settings → Privacy & Security** the first time you run them.

To build for your current OS only:

```bash
go build -ldflags="-s -w" -o mainupdate
```

## Requirements

- Go 1.25.4 or later
- Git repository with an `origin` remote and a detectable default branch

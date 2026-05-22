# mainupdate

A simple Go tool to update your local default branch from origin.

## What it does

1. Checks that your working tree is clean
2. Fetches all branches and tags from origin
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

## Requirements

- Go 1.25.4 or later
- Git repository with an `origin` remote and a detectable default branch

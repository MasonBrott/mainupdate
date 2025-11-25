# mainupdate

A simple Go tool to update your local main branch from origin.

## What it does

1. Checks that your working tree is clean
2. Checks out the main branch
3. Fetches all branches and tags from origin
4. Pulls the latest changes from origin/main

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
- Git repository with a `main` branch


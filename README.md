# Judge0 Backend

## Go Version

This project uses **Go 1.24.5**

Make sure Go 1.24.5 is installed before running or deploying the project.

## Prerequisites

- Go 1.24.5
- Vercel CLI

Install Vercel CLI if not already installed:

```bash
npm install -g vercel
```

## Local Development

For local development, run the main file from the `cmd` folder.

```bash
go run cmd/judge0-be/main.go
```

## Testing Before Commit

Before committing your changes, always test using Vercel local development mode.

```bash
vercel dev -d
```

This ensures the project works correctly in the Vercel environment.

## Notes

- Do not commit without testing locally
- Ensure Go version matches the required version

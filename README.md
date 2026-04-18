# create-pocketbase-template

A minimal GitHub template repository for starting a **Go + PocketBase** backend with Go-powered migrations wired from day one.

This template is intentionally manual and simple. You create a new repository from this template, initialize your Go module, update one import placeholder, and run the documented commands.

## What this template includes

- `main.go` – PocketBase app entrypoint with migration command registration.
- `migrations/migrations.go` – placeholder package so migrations are importable before your first snapshot.
- `.gitignore` – practical defaults for Go + PocketBase projects.

## Why there is no `go.mod` in this template

`go.mod` is intentionally **not** committed here.

After you generate your own repository from this template, you should create a module with your real repository path:

```bash
go mod init github.com/<your-username>/<your-new-repo>
```

This avoids committing a fake module path in the template and keeps setup explicit.

## Requirements

Install these before using the template:

- **Go** (recommended current stable)
- **Git**
- A **GitHub account**

Optional but useful:

- **GitHub CLI (`gh`)** for faster repo workflows

## 1) Create a new repo from this template

1. Open this template repository on GitHub.
2. Click **Use this template**.
3. Create your new repository (for example: `my-pocketbase-api`).

## 2) Clone your new repository

```bash
git clone https://github.com/<your-username>/<your-new-repo>.git
cd <your-new-repo>
```

## 3) Initialize the Go module

Run:

```bash
go mod init github.com/<your-username>/<your-new-repo>
```

Example:

```bash
go mod init github.com/acme/my-pocketbase-api
```

## 4) Update required placeholder in `main.go`

Current placeholder import:

```go
_ "your/module/path/migrations"
```

Replace it with your actual module path + `/migrations`, for example:

```go
_ "github.com/acme/my-pocketbase-api/migrations"
```

## 5) Install dependencies

After module init and import update:

```bash
go mod tidy
```

## 6) Run the app

```bash
go run . serve
```

PocketBase default local UI URL is typically `http://127.0.0.1:8090/_/`.

## 7) Create the initial PocketBase superuser

With the app running, open the admin UI in your browser and follow the first-time superuser creation flow.

## 8) Create your first migration snapshot

After collections are created/edited in the admin UI:

```bash
go run . migrate collections
```

This generates migration files under `migrations/`.

## 9) Recommended branch workflow (manual)

After your initial setup on `main`, run:

```bash
git add .
git commit -m "Initial setup from template"
git push origin main

git checkout -b test
git push -u origin test

git checkout -b dev
git push -u origin dev
```

Continue active development on `dev`.

---

## Notes

- This repository is the **template** repo (`create-pocketbase-template`).
- Future automation belongs in a separate repo (`create-pocketbase`).
- Keep this template focused on clean boilerplate + clear manual instructions.


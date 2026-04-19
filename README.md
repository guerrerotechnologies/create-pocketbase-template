# create-pocketbase-template

Minimal GitHub template for a fresh **Go + PocketBase** backend using PocketBase's official migration workflow.

## Purpose

Use this repo as a clean manual starting point:
- create a new repository from this template
- set your module/import placeholders
- run PocketBase
- generate migrations after schema changes

This template is intentionally minimal and does **not** include automation logic.

## What is included

- `main.go` (PocketBase app + migration command registration)
- `go.mod` / `go.sum` (starter dependencies; update module path)
- `.gitignore`
- `README.md`

## Important migration rule

In `main.go`, the migrations import stays commented at first:

```go
// _ "github.com/<your-username>/<your-repo>/migrations"
```

Do **not** uncomment it until your first migration file exists.

## Setup

### 1) Create a new repository from this template

1. Open this repository on GitHub.
2. Click **Use this template**.
3. Create your new repository.

### 2) Clone your new repository

```bash
git clone https://github.com/<your-username>/<your-repo>.git
cd <your-repo>
```

### 3) Update required placeholders

#### Update module path in `go.mod`

Replace:

```go
module github.com/guerrerotechnologies/create-pocketbase-template
```

with your real module path, for example:

```go
module github.com/acme/my-pocketbase-api
```

#### Update migrations import path in `main.go`

Keep it commented initially, but set the correct path:

```go
// _ "github.com/acme/my-pocketbase-api/migrations"
```

## Run commands

### Install/update deps

```bash
go mod tidy
```

### Run app

```bash
go run . serve
```

PocketBase admin UI is typically at `http://127.0.0.1:8090/_/`.

### Create first superuser

With the app running, open the admin UI and complete the first-time superuser flow.

## Official migration workflow

1. Run the app:

```bash
go run . serve
```

2. Create/edit collections in the PocketBase dashboard.
3. Stop the app.
4. Create the migrations folder manually:

```bash
mkdir migrations
```

5. Generate your first migration snapshot:

```bash
go run . migrate collections
```

6. After the first migration file exists, uncomment the migrations import in `main.go`:

```go
_ "github.com/<your-username>/<your-repo>/migrations"
```

7. Run the app again:

```bash
go run . serve
```

## Branch workflow for new repos

After your initial commit on `main`:

```bash
git checkout -b test
git push -u origin test

git checkout -b dev
git push -u origin dev
```

Continue development on `dev`.

---

## Scope note

- This repository: `create-pocketbase-template` (manual template only)
- Future separate automation repository: `create-pocketbase`

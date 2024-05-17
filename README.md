# RSS feed aggregator

Backend REST API server in golang, that allows aggregate data from RSS feeds and add different RSS feeds for podcast, blog, post etc that will automatically collect those posts from feeds with scrapper and download/save them into database so that we can view them later.

## Usage Instruction

### Requirements

- Golang > 1.20
- PostgreSQL > 14.0

### Run Localy Step-by-step

- Rename `.env.example` to `.env` and configure your environment variables
- To run server first time: `go build` after: `go build | .\rss-feed-aggr.exe` in powershell, in bash replace `|` to `&&`: `go build && .\rss-feed-aggr.exe` or just run builded executable file.
- Create database in PostgreSQL using PgAdmin or CLI.
- Change `DB_URL` in `.env` file with these structure: DB_URL=`postgres://username:password@localhost:5432/db_name?sslmode=disable`
- Install two CLI tools that allow working with ORM and Migrations
    - (additional) `go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest` (ORM) check with `sqlc version`
    - `go install github.com/pressly/goose/v3/cmd/goose@latest` (DB Migrations) check with `goose --version`
- To apply goose migrations:
    - `cd sql/schema`
    - up: `goose postgres postgres://username:password@localhost:5432/db_name up`
    - down: `goose postgres postgres://username:password@localhost:5432/db_name down`
- (additional) sqlc Generator ORM (SQL query in `sql/queries/` to Golang code to `internal/database`):
    - `sqlc generate`
- 

### Routers/Endpoints

## Tech stack

- Golang
- PostgreSQL

### Dependency Third-party Packages/Modules

- `github.com/go-chi/chi` REST API router/endpoint HTTP maker lightweight service [docs](https://go-chi.io/#/README)
- `github.com/go-chi/cors` router CORS middleware rules
- `github.com/joho/godotenv` load `.env` file as environment variable
- `github.com/lib/pq` postgresql driver
- `github.com/pressly/goose` CLI tool for Migrations with Raw SQL [docs](https://pressly.github.io/goose/)
- `github.com/sqlc-dev/sqlc` CLI tool for Generate Golang Code with Raw SQL [docs](https://docs.sqlc.dev/en/stable/index.html)
- `github.com/google/uuid` UUID generator

## TODO

- [ ] Add Docker, makefile
- [ ] OpenAPI, Swagger API docs
- [x] Connect PostgreSQL database, and ORM(generating go code sqlc)/Migrations (goose) future with raw SQL query
- [x] Authentication w/ API keys
- [x] Aggregation scraper/worker

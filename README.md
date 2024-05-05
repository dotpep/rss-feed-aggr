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
- Change `DB_URL` in `.env` file with these structure: DB_URL=`postgres://username:password@localhost:5432/db_name`
- Install two CLI tools that allow working with ORM and Migrations
    - `go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest` (ORM) check with `sqlc version`
    - `go install github.com/pressly/goose/v3/cmd/goose@latest` (DB Migrations) check with `goose --version`
- To apply goose migrations:
    - `cd sql/schema`
    - up: `goose postgres postgres://username:password@localhost:5432/db_name up`
    - down: `goose postgres postgres://username:password@localhost:5432/db_name down`
- 

### Routers/Endpoints

## Tech stack

- Golang
- PostgreSQL

### Dependency Third-party Packages/Modules

## TODO

- [ ] Add Docker, makefile
- [ ] Connect PostgreSQL database, and ORM/Migrations future

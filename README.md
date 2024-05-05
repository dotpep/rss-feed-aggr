# RSS feed aggregator

Backend REST API server in golang, that allows aggregate data from RSS feeds and add different RSS feeds for podcast, blog, post etc that will automatically collect those posts from feeds with scrapper and download/save them into database so that we can view them later.

## Usage Instruction

- Rename `.env.example` to `.env` and configure your environment variables
- To run server first time: `go build` after: `go build | .\rss-feed-aggr.exe` in powershell, in bash replace `|` to `&&`: `go build && .\rss-feed-aggr.exe` or just run builded executable file.

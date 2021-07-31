# A simple job fetcher written in Go

This is a simple job crawler cli tool. Extracts data from web pages and allows to search by keyword.

Follow these steps to try the crawler out:

- Run the image `docker run -p 7317:7317 ghcr.io/go-rod/rod` to manage Rod launcher. [Rod](https://go-rod.github.io) is a tool that allows us to scrape client-side rendered pages.
- Generate the executable `go build -o go-fetcher`, you can specify a different output name if you wish
- `./go-fetcher <parser-id> <keyword>` - this command will extract jobs based on the parser rules, search for the keyword and print urls of jobs that contain that keyword 

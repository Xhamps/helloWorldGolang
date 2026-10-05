# htmx-hello

Minimal Go + htmx "Hello, World" web app. This is a **test project**: its purpose
is to serve as a small, dependency-free workload for checking infrastructure
(builds, container images, deploy pipelines, health checks, routing, etc.), not
to be a real application.

## What it does

A single Go binary serves two routes on port `8080`:

| Route    | Response                                                        |
|----------|-----------------------------------------------------------------|
| `/`      | Full HTML page (embedded `index.html`) with an "Initialize" button |
| `/hello` | HTML fragment with `HELLO, WORLD` and the current server time    |

Clicking the button makes htmx fetch `/hello` and swap the fragment into the page.
The template is embedded in the binary with `embed.FS`, so the build output is a
single self-contained executable. The only runtime dependency is the Go standard
library; htmx is loaded from a CDN in the browser.

## Requirements

- Go 1.26+

## Usage

```sh
make run      # start the server on http://localhost:8080
make build    # compile ./htmx-hello
make dev      # rebuild + restart on file change
make test     # go test ./...
make vet      # go vet ./...
make fmt      # gofmt
make clean    # remove the binary
```

Or without make:

```sh
go run .
```

## Smoke check

```sh
curl -s localhost:8080/        # expect HTML page, HTTP 200
curl -s localhost:8080/hello   # expect "<h1 ...>HELLO, WORLD</h1>..."
curl -s -o /dev/null -w '%{http_code}\n' localhost:8080/nope   # expect 404
```

## Layout

```
main.go     HTTP server and routes
index.html  page template (embedded at build time)
Makefile    dev/build helpers
```

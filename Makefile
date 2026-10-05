BIN := htmx-hello

.PHONY: run build test fmt vet tidy clean dev

run: ## start the server
	go run .

build: ## compile binary
	go build -o $(BIN) .

test:
	go test ./...

fmt:
	gofmt -w .

vet:
	go vet ./...

tidy:
	go mod tidy

clean:
	rm -f $(BIN)

# ponytail: polling reload via `find`, swap for air if the file list grows
dev: ## rebuild+restart on change
	@while true; do \
		go run . & PID=$$!; \
		STAMP=$$(stat -f %m main.go index.html go.mod | sort -n | tail -1); \
		while [ "$$(stat -f %m main.go index.html go.mod | sort -n | tail -1)" = "$$STAMP" ]; do sleep 1; done; \
		kill $$PID 2>/dev/null; wait $$PID 2>/dev/null; \
	done

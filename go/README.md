# <lang> Producer/Consumer

Kafka Producer/Consumer code in GO

NOTES:
- Serialization for payload: JSON -> Bytes
- Deserialization for payload: Bytes -> JSON

# Requirements

- GO (go1.22.5)

- [Install GO](https://go.dev/doc/install)
- [Upgrade GO in Mac OS](https://stackoverflow.com/questions/42952979/go-version-command-shows-old-version-number-after-update-to-1-8)

## Dependencies

```bash
go mod tidy
```

# Running application

## Run application:

```bash
go run .
```

## Build binary
```bash
go build
```

## Install binary in `go` path:

List path directory where go app will be installed
```bash
go list -f '{{.Target}}'
```

Install application
```bash
go install
```

# Running Consumer

```bash
cd consumer
go run .
```

# Running Producer

```bash
cd producer
go run .
```

# Run tests:

For Consumer and Producer, just go into each directory and run the `go test` command

```bash
go test

# OR
go test -v
```

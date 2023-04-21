# mini-api

Mini Admin API

## Build

### Example version 1.0.0

```shell
go build -ldflags="-X mini-api/config.Version=1.0.0"
```

## Run

### Example

```shell
go run main.go --help
```

## Installation

### init data

```text
1. change env file INIT_DATA=true 
2. run make api
```

### Prerequisite

- [Swaggo](https://github.com/swaggo/swag)

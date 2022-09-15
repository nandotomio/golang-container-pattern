# Project: golang-container-pattern

A simple API in GO using container pattern

## Requirements

This project requires the following dependencies to run:

- Golang v1.14+
- docker v20.10+
- docker compose v2.4+
- GNU Make v3.81+

## Getting started (development)

```bash
# to start development run the following commands
make dbup # start postgres database
make dev # start server

# after you finish development, you need to manually turn down the database container with
make dbdown
```
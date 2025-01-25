# Go API template

This is a template for a Go API project. It includes the following features:

- Clean Architecture
- Swagger Documentation
- Dockerized
- Makefile for easy setup
- API Security
- API Logging
- API Testing

## Requirements

- [Git](https://git-scm.com/downloads) - Git CLI
- [Code Editor](https://code.visualstudio.com) - Basic code editor
- [Golang](https://go.dev) - Golang
- [Docker](https://docs.docker.com/get-docker/) - For containerization and local environment (if you're not using Makefile)
- [GoSec](https://github.com/securego/gosec) & [Golang-Lint](https://github.com/golangci/golangci-lint) - Go Validators
- [GoFumpt](https://github.com/mvdan/gofumpt) - Stricter GoFMT
- [Redoc-CLI](https://redocly.com/docs/redoc/deployment/cli) - For generating API Documentation (npm install -g @redocly/cli)

## Getting started

### Native

Download all necessary modules and run the application

```
go mod download
make run
```

### Docker

#### First application start

```
make docker
```

#### After the first start

After first setup you can use docker-compose to start and stop / restart the application

```
docker-compose up
docker-compose down
```

## Swagger

Swagger documentation is available at [http://localhost:8080/docs](http://localhost:8080/docs)

To generate the Swagger documentation, you can run the following command:

```
make swagger
```

## Contributing

You can find the detailed contributing guidelines by navigating to the [`CONTRIBUTE.md`](/CONTRIBUTE.md/) file in the repository.

## Docs

- [Swagger](http://localhost:8080/docs) - Swagger Documentation, available after running the application

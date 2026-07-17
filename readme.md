# Calculator Backend

![Coverage](https://img.shields.io/badge/coverage-0%25-lightgrey)

Coverage badge is published from GitHub Pages via the CI workflow.

## Overview
A production-style Go microservice that exposes a single HTTP endpoint for evaluating mathematical expressions.

## Architecture
The service follows a layered design:
- API layer for HTTP handling
- Middleware for logging and recovery
- Validator for request validation
- Service layer for orchestration
- Evaluator for expression computation

## Local Setup

### Prerequisites
- Go 1.24 or newer
- Docker and Docker Compose

### 1. Clone and install dependencies
```bash
git clone <repository-url>
cd Calculator-backend
go mod download
```

### 2. Run locally with Go
```bash
make run
```

The service will start on port 8080 by default.

Example requests:
```bash
curl http://localhost:8080/health
curl -X POST http://localhost:8080/calculate -H "Content-Type: application/json" -d '{"expression":"2+2*2"}'
```

### 3. Run locally with Docker
```bash
docker compose up --build -d
```

The Docker setup exposes the app on port 8081 on the host:
```bash
curl http://localhost:8081/health
curl -X POST http://localhost:8081/calculate -H "Content-Type: application/json" -d '{"expression":"2+2*2"}'
```

## Test Guide

### Run the full test suite
```bash
make test
```

### Run tests with coverage
```bash
go test ./tests ./cmd/server -coverprofile=coverage.out -coverpkg=calculator-backend/internal/api,calculator-backend/internal/config,calculator-backend/internal/evaluator,calculator-backend/internal/middleware,calculator-backend/internal/service,calculator-backend/internal/validator,calculator-backend/cmd/server
go tool cover -func=coverage.out
```

### Lint the code
```bash
make lint
```

### Format the code
```bash
make fmt
```

## API Contract

### POST /calculate
Evaluate a mathematical expression.

Request body:
```json
{
  "expression": "2+2*2"
}
```

Success response:
```json
{
  "result": 6
}
```

Error response:
```json
{
  "error": "division by zero"
}
```

### GET /health
Check that the service is running.

Success response:
```json
{
  "status": "ok"
}
```

## Folder Structure
- cmd/server: application entrypoint
- internal/api: router and handlers
- internal/middleware: logging and panic recovery
- internal/validator: request validation
- internal/service: business coordination
- internal/evaluator: expression evaluation
- internal/models: DTOs
- internal/config: environment configuration
- tests: unit tests

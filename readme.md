# Calculator Backend

## Overview
A production-style Go microservice that exposes a single HTTP endpoint for evaluating mathematical expressions.

## Architecture
The service follows a layered design:
- API layer for HTTP handling
- Middleware for logging and recovery
- Validator for request validation
- Service layer for orchestration
- Evaluator for expression computation

Coverage is collected in CI and uploaded as a workflow artifact. Download the artifact from the Actions run to inspect the generated report locally.

> To review the coverage report, open the Actions run, download the coverage artifact, and inspect the generated HTML or text report from the downloaded files.

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

### 3. Run without Docker
If Docker is not available, you can run the services locally.

Start the backend first:
```bash
cd Calculator-backend
go run ./cmd/server
```

The backend will be available at http://localhost:8081.

Then start the frontend in a second terminal:
```bash
cd frontend
npm install
VITE_API_URL=http://localhost:8081 npm run dev
```

The frontend will be available at http://localhost:5173 once the backend is up.

### 4. Run locally with Docker
```bash
docker compose up --build -d
```

This starts both services together:
- Frontend: http://localhost:5173
- Backend: http://localhost:8081

You can verify the backend with:
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

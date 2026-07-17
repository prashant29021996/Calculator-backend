# Calculator Frontend

A responsive React + TypeScript calculator frontend for the calculator backend. The UI validates expressions locally, sends only valid requests to the backend, and keeps a short calculation history for convenience.

## Features

- Responsive layout for desktop and mobile devices
- Accessible form controls, keyboard support, visible focus states, and screen-reader-friendly status updates
- Local validation before API calls
- Result caching to avoid redundant backend requests for repeated expressions
- Production-ready build output served by Nginx in Docker

## Tech stack

- React 18
- TypeScript
- Vite
- Vitest + Testing Library
- Docker / Nginx

## Development

Install dependencies:

```bash
cd frontend
npm install
```

Run locally:

```bash
npm run dev
```

Run tests:

```bash
npm test
```

Build for production:

```bash
npm run build
```

## Environment variables

- VITE_API_URL: Backend base URL used by the frontend. Defaults to http://localhost:8081.

## Usage examples

The calculator supports common arithmetic expressions and a few convenience operators:

- Exponentiation: enter `5^2` and it will be converted to the backend-safe form before submission.
- Square root: enter `sqrt(9)` to evaluate the square root of 9.
- Percentages: enter expressions such as `25%30` to compute 25% of 30, or `10%` for 10/100.

## Docker

The frontend image is built in two stages:

1. Build the Vite app.
2. Serve the built static files with Nginx.

To run the full stack with Docker Compose from the repository root:

```bash
docker compose up --build -d
```

This brings up both services:
- Frontend: http://localhost:5173
- Backend: http://localhost:8081

You can also build the frontend image directly:

```bash
docker build -t calculator-frontend .
```

## Accessibility and responsiveness

The interface is designed to be accessibility compliant and responsive. It uses semantic buttons, labeled inputs, live regions for status updates, visible focus styles, and a layout that adapts cleanly across desktop and mobile screens.

## Notes for future improvements

The current frontend is already functional and tested, but future work could include:

- a more sophisticated component architecture for additional calculator modes
- richer visual theming and animation polish
- broader internationalization and localization support

.PHONY: run test lint fmt docker coverage clean

run:
	go run ./cmd/server

test:
	go test ./...

lint:
	@set -e; \
	GOBIN="$$(go env GOPATH)/bin"; \
	if [ ! -x "$$GOBIN/golangci-lint" ]; then \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b "$$GOBIN" v1.64.8; \
	fi; \
	"$$GOBIN/golangci-lint" run

fmt:
	go fmt ./...

docker:
	docker build -t calculator-backend .

coverage:
	go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out

clean:
	rm -f coverage.out

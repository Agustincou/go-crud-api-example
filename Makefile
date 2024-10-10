.PHONY: lint

lint:
	golangci-lint run && go run ./pkg/custom_linter/pkgcheck.go -- ./...

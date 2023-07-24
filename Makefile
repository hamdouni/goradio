ifndef VERBOSE
.SILENT:
endif

## audit: run quality control checks
.PHONY: audit
audit:
	go mod verify
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3 run --max-same-issues 0 --max-issues-per-linter 0

.PHONY: test
test:
	go test -v -race -buildvcs -count=1 ./... 


.PHONY: ci-check fmt-check vet test

ci-check: fmt-check vet test

fmt-check:
	@test -z "$$(gofmt -l .)" || { gofmt -l .; exit 1; }

vet:
	go vet ./...

test:
	go test -race -count=1 ./...

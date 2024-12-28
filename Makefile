.PHONY: gen fmt

gen:
	go run testdata/gen.go

fmt:
	go mod tidy
	gofmt -s -w .
	goarrange run -r .
	golangci-lint run ./...

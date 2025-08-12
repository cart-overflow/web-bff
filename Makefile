
dev-setup:
	lefthook install

fmt:
	golangci-lint fmt

lint:
	golangci-lint run
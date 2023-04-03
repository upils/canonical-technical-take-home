.PHONY: build
build:
	CGO_ENABLED=0 go build -o build/ .

.PHONY: test
test:
	go test 

.PHONY: run
run:
	go run main.go -path=tests/test.txt -n=5
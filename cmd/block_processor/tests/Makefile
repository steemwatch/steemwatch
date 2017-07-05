.PHONY: build

build:
	go test -c -o build/tests

docker/exe:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go test -c -o build/tests_entrypoint

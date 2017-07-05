.PHONY: build

build:
	go build -o build/steemd

docker/exe:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o build/steemd_entrypoint
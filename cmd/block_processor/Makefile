.PHONY: build

build:
	go build -o build/block_processor

test:
	cd tests && ./run.sh

docker/exe:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o build/block_processor_entrypoint

docker/build: docker/exe
	docker build --tag steemwatch/block_processor .

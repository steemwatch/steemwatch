.PHONY: build

build:
	go build -o build/event_emitter

test:
	cd tests && ./run.sh

docker/exe:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o build/event_emitter_entrypoint

docker/build: docker/exe
	docker build --tag steemwatch/event_emitter .

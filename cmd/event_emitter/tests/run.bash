#!/usr/bin/env bash

set -e

PROJECT='block_processor_tests'
COMPOSE="docker-compose --project-name $PROJECT"

# Build tests.
function build_tests() {
	echo
	echo '---> Build the tests Docker image'
	make docker/exe
	$COMPOSE build test_runner
}

# Build steemd.
function build_steemd() {
	echo
	echo '---> Build the steemd mock Docker image'
	make -C steemd docker/exe
	$COMPOSE build steemd
}

# Build block_processor.
function build_block_processor() {
	echo
	echo '---> Build the block_processor Docker image'
	make -C .. docker/exe
	$COMPOSE build block_processor
}

# Register a cleanup function to remove used containers.
function remove_containers() {
	test -n "$KEEP_CONTAINERS" && return

	echo
	echo '---> Remove used Docker containers'
	$COMPOSE stop
	$COMPOSE rm -v -f
}

# Create containers.
function create_containers() {
	echo
	echo '---> Create Docker containers'
	$COMPOSE create
}

# Run tests.
function run() {
	echo
	echo '---> Run tests'
	$COMPOSE run --rm test_runner \
		/tests_entrypoint -test.v -test.timeout 30s $TEST_FLAGS
}

if [ -z "$SKIP_BUILD" ]; then
	build_tests
    build_steemd
	build_block_processor
fi

trap remove_containers EXIT
create_containers
run

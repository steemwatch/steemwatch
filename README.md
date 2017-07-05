# SteemWatch 2.0 Backend

This monorepo contains all sources for SteemWatch 2.0 backend.

## Status

SteemWatch 2.0 is happening right now. Anything is a subject to change,
anything may and everything will change. So in case you want to use a package
from this repository, make sure to lock the revision or just copy the sources out,
with the associated copyright notice as required by Apache License 2.0, obviously.

## Repository Organization

`cmd/` subdirectory contains all executable packages.
The shared packages can be found in `pkg/`.

## Build Process

We are using [dep](https://github.com/golang/dep), so you need to fetch
the dependencies before you can start doing anything meaningful.

Once the dependencies are there, you can `cd` to the executable package
directory of your choice. The package can be build using `go build` or
just by invoking `make`. Running `make docker/build` builds the service
Docker image where it makes sense.

Occasionally there is also `make test` available for running tests.

## License

Apache License 2.0, see the `LICENSE` file.
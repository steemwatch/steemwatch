# pb

This directory contains various Protocol Buffers definitions used by SteemWatch.

To generate Go sources, invoke `protoc` in the relevant subdirectory:

```
export PATH="$PATH/$GOPATH/bin"
go get -u github.com/gogo/protobuf/protoc-gen-gogo
protoc --proto_path="$GOPATH/src" --proto_path=. --gogo_out=. *.proto
```
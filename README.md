Requirement:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```


To generate proto go code:

```
protoc --go_out=. --go-grpc_out=. proto/pokemon.proto
```

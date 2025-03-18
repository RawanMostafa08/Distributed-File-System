```
go mod tidy
```

```
go run grpc\master\main.go
```

```
protoc --go_out=. --go-grpc_out=. upload.proto
```
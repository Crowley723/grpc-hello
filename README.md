# GRPC Hello

This is a simple (ai-generated) go application that is meant to show how gRPC works in go.

The application is a simple greet/goodbye system. The server listens for connections and prints messages sent by clients. The client connects, sends a hello message (with an optional name), and then sends a goodbye message.

The only code that is not generated from the `greet.proto` file is in `main.go`.

Try regenerating the code in the `greet` package yourself:
1. Install the protobuf compiler (apt: `protobuf-compiler`)
2. Install the required plugins for go: `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
and `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`
3. Make sure they are in your path: `export PATH="$PATH:$(go env GOPATH)/bin"`
4. Delete the files in `greet/`
5. Generate the files: `protoc --go_out=./greet --go_opt=paths=source_relative --go-grpc_out=./greet --go-grpc_opt=paths=source_relative greet.proto`

This will generate `greet/greet.pb.go` and `greet/greet_grpc.pb.go` which you notice will allow the methods inside them to be imported and called from `main.go`.

Some very cool stuff.
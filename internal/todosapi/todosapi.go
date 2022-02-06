package todosapi

//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen -generate types -package todosapi -o todosapi-types.gen.go ../../openapi/todo.yaml
//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen -generate server,spec -package todosapi -o todosapi-server.gen.go ../../openapi/todo.yaml
//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen -generate client -package todosapi -o todosapi-client.gen.go ../../openapi/todo.yaml
//go:generate gofmt -s -w todosapi-server.gen.go todosapi-types.gen.go todosapi-client.gen.go

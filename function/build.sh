GOARCH=wasm GOOS=js go build -o function.wasm function.go
go build -o server server.go
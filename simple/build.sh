GOARCH=wasm GOOS=js go build -o simple.wasm simple.go
go build -o server server.go
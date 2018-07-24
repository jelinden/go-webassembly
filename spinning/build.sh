GOARCH=wasm GOOS=js go build -o spinning.wasm spinning.go
go build -o server server.go
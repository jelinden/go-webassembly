GOARCH=wasm GOOS=js go build -o valuetoelement.wasm valuetoelement.go
go build -o server server.go
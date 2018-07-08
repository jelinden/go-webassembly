GOARCH=wasm GOOS=js go build -o redbox.wasm redbox.go
go build -o server server.go
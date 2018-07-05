package main

import (
	"syscall/js"
)

func main() {
	doc := js.Global().Get("document")
	doc.Call("getElementById", "hello").Set("innerHTML", "Hello Wasm!")
}

package main

import (
	"syscall/js"
)

const width = "width:200px;"
const height = "height:200px;"
const backgroundColor = "background-color:red;"
const center = `margin:auto;`
const font = "color:white;"

func main() {
	doc := js.Global().Get("document")
	element := doc.Call("getElementById", "redbox")
	element.Set("innerHTML", "Hello Wasm!")
	element.Call("setAttribute", "style", width+height+backgroundColor+center+font)
}

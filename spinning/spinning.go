package main

import (
	"fmt"
	"syscall/js"
)

const width = "width:200px;"
const height = "height:200px;"
const backgroundColor = "background-color:red;"
const center = `margin:auto;`
const font = "color:white;"

var element js.Value
var doc js.Value
var counter int
var beforeUnloadChannel = make(chan struct{})

func main() {
	doc = js.Global().Get("document")
	element = doc.Call("getElementById", "spinning")
	element.Call("setAttribute", "style", width+height+backgroundColor+center+font)

	cb := js.NewCallback(spin)
	defer cb.Release()
	runSpin := js.Global().Get("runSpin")
	runSpin.Invoke(cb)
	<-beforeUnloadChannel
}

func spin(args []js.Value) {
	element.Call("setAttribute", "style", rotate(counter*2)+width+height+backgroundColor+center+font)
	counter++
	fmt.Println("spin", counter)
}

func rotate(degree int) string {
	return fmt.Sprintf(`-webkit-transition: -webkit-transform 1.5s linear;transform: rotate(%vdeg);`, degree)
}

func beforeUnload(event js.Value) {
	beforeUnloadChannel <- struct{}{}
}

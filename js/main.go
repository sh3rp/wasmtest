// +build wasm
package main

import "syscall/js"

type jsObject = js.Value

var document jsObject

func main() {
	canvas := js.Global().Get("document").Call("getElementById", "viewport")
	ctx := canvas.Call("getContext", "2d")
	ctx.Set("font", "30px Arial")
	ctx.Call("fillText", "Hello world!", 10, 50)
}

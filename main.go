//go:build js && wasm

package main

import (
	"fmt"
	"syscall/js"
	"time"
)

var (
	counter int
	clock   string
)

func updateClock() {
	for {
		clock = time.Now().Format("15:04:05.000") // Format jam hingga milidetik
		js.Global().Get("document").Call("getElementById", "clock").Set("innerHTML", clock)
		time.Sleep(200 * time.Millisecond) // Update setiap 100 milidetik
	}
}

func increment(this js.Value, p []js.Value) interface{} {
	counter++
	updateDisplay()
	return nil
}

func decrement(this js.Value, p []js.Value) interface{} {
	counter--
	updateDisplay()
	return nil
}

func updateDisplay() {
	js.Global().Get("document").Call("getElementById", "counter").Set("innerHTML", fmt.Sprintf("%d", counter))
}

func main() {
	// Inisialisasi nilai
	counter = 0
	updateDisplay()

	// Mendaftarkan fungsi JavaScript
	js.Global().Set("increment", js.FuncOf(increment))
	js.Global().Set("decrement", js.FuncOf(decrement))

	// Mulai thread untuk mengupdate jam
	go updateClock()

	// Prevent the program from exiting
	select {}
}

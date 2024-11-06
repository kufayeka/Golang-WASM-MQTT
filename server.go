package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Setup Gin router
	r := gin.Default()
	r.StaticFile("/wasm_exec.js", "./wasm_exec.js") // Serve wasm_exec.js file
	r.StaticFile("/main.wasm", "./main.wasm")       // Serve .wasm file
	r.LoadHTMLFiles("index.html")                   // Serve HTML file

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// Jalankan server
	err := r.Run(":8888")
	if err != nil {
		return
	}
}

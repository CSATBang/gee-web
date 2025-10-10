package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.Default()

	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello BastienHyde\n")
	})

	r.GET("/panic", func(c *gee.Context) {
		name := []string{"BastienHyde"}
		c.String(http.StatusOK, name[100])
	})
	r.Run(":10020")
}

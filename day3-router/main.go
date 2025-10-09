package main

import (
	"gee"
	"net/http"
)

/*
可以用Postman进行测试
也可以用curl工具
$ curl "http://localhost:10020/hello/BastienHyde"
hello BastienHyde, you're at /hello/BastienHyde

$ curl "http://localhost:10020/assets/css/BastienHyde.css"
{"filepath":"css/BastienHyde.css"}
*/
func main() {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.GET("/hello", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s,you`re at %s\n", c.Query("name"), c.Path)
	})
	r.GET("hello/:name", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})
	r.GET("/assets/*filepath", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"filepath": c.Param("filepath"),
		})
	})
	r.Run(":10020")
}

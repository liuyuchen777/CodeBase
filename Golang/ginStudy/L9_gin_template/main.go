/*
 * @Author: Liu Yuchen
 * @Date: 2021-03-31 15:37:38
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-03-31 17:09:09
 * @Description: Say Something, Motherfuncker
 * @FilePath: /go/src/goStudy/gin_demo/L9_Gin_template/main.go
 * @GitHub: https://github.com/liuyuchen777
 */

package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func indexDemo(c *gin.Context) {
	// http请求状态码
	c.HTML(http.StatusOK, "main/index.tmpl", gin.H{
		"title": "http://liuyuchen777.github.io", // rendor
	})
}

// 静态文件
// .css js pic

func main() {
	r := gin.Default()
	// add self define function
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	// r.LoadHTMLFiles("templates/posts/index.tmpl", "templates/users/index.tmpl") // parse
	// load static file
	r.Static("/xxx", "./statics")
	// load template
	r.LoadHTMLGlob("./templates/**/*")
	// XZr.LoadHTMLFiles("./templates/index.tmpl")
	// 可以多次加载
	r.GET("/main/index", indexDemo)
	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "I am from posts/index.html",
		})
	})

	r.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "<a href='http://liuyuchen777.github.io'>LiuyuchenBlog</a>",
		})
	})

	r.Run(":9090")
}

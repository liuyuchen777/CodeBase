/*
 * @Author: Liu Yuchen
 * @Date: 2021-03-31 17:12:48
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-03-31 17:32:51
 * @Description: Say Something, Motherfuncker
 * @FilePath: /go/src/goStudy/gin_demo/L10_template_web/main.go
 * @GitHub: https://github.com/liuyuchen777
 */

package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.Run(":9090")
}

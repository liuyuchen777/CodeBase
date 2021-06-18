/*
 * @Author: Liu Yuchen
 * @Date: 2021-04-01 09:04:14
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-04-01 09:33:31
 * @Description: Say Something, Motherfuncker
 * @FilePath: /go/src/goStudy/ginStudy/L13_form/main.go
 * @GitHub: https://github.com/liuyuchen777
 */

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLFiles("./login.html", "./index.html")

	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	router.POST("/login", func(c *gin.Context) {
		// get form
		// username := c.PostForm("username")
		// password := c.PostForm("password")

		// username := c.DefaultPostForm("username", "somebody")
		// password := c.DefaultPostForm("xxx", "***")

		username, ok := c.GetPostForm("username")
		if !ok {
			username = "lyc"
		}
		password, ok := c.GetPostForm("password")
		if !ok {
			password = "***"
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"username": username,
			"password": password,
		})
	})

	router.Run(":9090")
}

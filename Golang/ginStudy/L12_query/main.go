/*
 * @Author: Liu Yuchen
 * @Date: 2021-03-31 18:03:17
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-03-31 19:28:24
 * @Description: gin to get query string
 * @FilePath: /go/src/goStudy/ginStudy/L12_query/main.go
 * @GitHub: https://github.com/liuyuchen777
 */

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/web", func(c *gin.Context) {
		// 获取request中的query string
		// GET请求?后面的事query string
		name := c.Query("query")
		age := c.Query("age")
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})
	router.Run(":9090")
}

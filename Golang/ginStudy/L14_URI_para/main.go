/*
 * @Author: Liu Yuchen
 * @Date: 2021-04-01 09:35:32
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-04-01 09:45:22
 * @Description: Say Something, Motherfuncker
 * @FilePath: /go/src/goStudy/ginStudy/L14_URI_para/main.go
 * @GitHub: https://github.com/liuyuchen777
 */
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "hello",
		})
	})

	router.GET("/user/:name/:age", func(c *gin.Context) {
		Name := c.Param("name")
		Age := c.Param("age") // string type
		c.JSON(http.StatusOK, gin.H{
			"username": Name,
			"age":      Age,
		})
	})

	router.GET("/blog/:year/:month", func(c *gin.Context) {
		Year := c.Param("year")
		Month := c.Param("month")
		c.JSON(http.StatusOK, gin.H{
			"year":  Year,
			"month": Month,
		})
	})

	router.Run(":9090")
}

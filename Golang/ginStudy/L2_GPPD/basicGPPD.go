/*
 * @Author: Liu Yuchen
 * @Date: 2021-03-31 09:39:05
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-03-30 10:47:49
 * @Description: Say Something, Motherfuncker
 * @FilePath: /Local_Lab/goStudy/gin_demo/basicGPPD.go
 * @GitHub: https://github.com/liuyuchen777
 */

package main

import "github.com/gin-gonic/gin"

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello golang",
	})
}

func main() {
	// create a default router engine
	r := gin.Default() // return default router engine

	r.GET("/hello", sayHello)
	// start service
	// RESTful API
	// 向前端返回JSON
	r.GET("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "GET",
		})
	})
	r.POST("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "POST",
		})
	})
	r.PUT("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "PUT",
		})
	})
	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "DELETE",
		})
	})
	r.Run()
}

/*
 * @Author: Liu Yuchen
 * @Date: 2021-04-02 00:59:37
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-04-02 01:32:21
 * @Description: Say Something, Motherfuncker
 * @FilePath: /go/src/goStudy/ginStudy/L17_redirect_routeG/main.go
 * @GitHub: https://github.com/liuyuchen777
 */

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/index", func(c *gin.Context) {
		/*
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		*/
		c.Redirect(http.StatusMovedPermanently, "http://www.google.com")
	})

	router.GET("/a", func(c *gin.Context) {
		// 跳转到/b对应路由处理函数
		c.Request.URL.Path = "/b" // 把请求的URL修改
		router.HandleContext(c)   // 继续后续的处理
	})

	router.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "b",
		})
	})

	// 路由组
	// 访问index的get请求
	router.GET("/any", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "GET",
		})
	})
	router.POST("/any", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "POST",
		})
	})
	router.DELETE("/any", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "DELETE",
		})
	})
	router.PUT("/any", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "PUT",
		})
	})
	// 处理所有的请求方法
	router.Any("/user", func(c *gin.Context) {
		switch c.Request.Method {
		case http.MethodGet:
			c.JSON(http.StatusOK, gin.H{
				"method": "GET",
			})
		case http.MethodPost:
			c.JSON(http.StatusOK, gin.H{
				"method": "GET",
			})
		case http.MethodDelete:
			c.JSON(http.StatusOK, gin.H{
				"method": "DELET",
			})
		case http.MethodPut:
			c.JSON(http.StatusOK, gin.H{
				"method": "PUT",
			})
		default:
			c.JSON(http.StatusOK, gin.H{
				"method": "AnyRequest",
			})
		}
	})
	// noRouter
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "liuyuchen",
		})
	})
	// route group
	// 视频的首页和详情页
	/* // not very effficient
	router.GET("/video/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "/video/index",
		})
	})
	*/
	videoGroup := router.Group("/video") // 公用的路由提取出来
	{
		videoGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "/video/index",
			})
		})
		videoGroup.GET("/xxx", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "/video/xxx",
			})
		})
		videoGroup.GET("/ooo", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "/video/ooo",
			})
		})
	}
	// 商城的首页和详情页
	router.GET("/shop/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "/shop/index",
		})
	})
	// 路由组支持嵌套

	router.Run(":9090")
}

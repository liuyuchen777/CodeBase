/*
 * @Author: Liu Yuchen
 * @Date: 2021-04-02 01:34:43
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-04-02 02:20:45
 * @Description: middleware
 * @FilePath: /go/src/goStudy/ginStudy/L18_middleware/main.go
 * @GitHub: https://github.com/liuyuchen777
 */
package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

/*
中间件适合处理一些公共的业务逻辑，eg：登录认证，权限校验，数据分页，记录日志，耗时统计
*/

func mainHandler(c *gin.Context) {
	fmt.Println("...mainHandler...")
	c.JSON(http.StatusOK, gin.H{
		"msg": "main",
	})
}

// 中间件函数：统计处理函数耗时
func m1(c *gin.Context) {
	fmt.Println("....m1 in...")
	start := time.Now()
	// go funcXX(c.Copy()) // 在funcXX中只能使用c的拷贝
	c.Next() // 调用后续的处理函数
	// c.Abort() // 阻止调用后续函数
	cost := time.Since(start)
	fmt.Printf("cost: %v\n", cost)
	fmt.Println("...m1 out...")
}

func m2(c *gin.Context) {
	fmt.Println("....m2 in...")
	c.Abort()
	c.Set("name", "lyc") // 跨中间件存值
	// return // 立即结束m2函数
	/*  if not c.Abort()
	....m1 in...
	....m2 in...
	...m2 out...
	...shopHandler...
	cost: 128.6µs
	...m1 out...
	*/
	fmt.Println("...m2 out...")
}

func authMiddleware(c *gin.Context) {
	// 是否登陆的判断
	// if 是登录用户
	// c.Next()
	// else
	// c.Abort()
}

// 一个判断用户是否登录的函数
func authMiddleware2(doCheck bool) gin.HandlerFunc {
	// 一些其他准备工作，eg：连接数据库等
	return func(c *gin.Context) {
		if doCheck {
			// 是否登陆的判断
			// if 是登录用户
			// c.Next()
			// else
			// c.Abort()
		} else {
			c.Next()
		}
	}
}

func main() {
	r := gin.Default() // 默认使用Logger()和Recovery()中间件
	// r := gin.New()
	// func (group *RouterGroup) GET(relativePath string, handlers ...HandlerFunc)
	// 顺序：m1, mainHandler

	r.Use(m1, m2, authMiddleware2(true)) // 全局注册中间件函数m1
	// 一部分路由使用用路由组

	r.GET("/main", mainHandler)
	r.GET("/shop", func(c *gin.Context) {
		name, _ := c.Get("name")
		fmt.Println("...shopHandler...")
		c.JSON(http.StatusOK, gin.H{
			"msg": name,
		})
	})

	r.Run(":9090")
}

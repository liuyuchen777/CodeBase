/*
 * @Author: Liu Yuchen
 * @Date: 2021-04-01 09:45:48
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-04-01 10:08:06
 * @Description: Say Something, Motherfuncker
 * @FilePath: /go/src/goStudy/ginStudy/L15_bind/main.go
 * @GitHub: https://github.com/liuyuchen777
 */

package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

// 参数绑定，query string或form中的信息与后端的结构体进行绑定

func main() {
	router := gin.Default()

	router.GET("/user", func(c *gin.Context) {
		/*
			username := c.Query("username")
			password := c.Query("password")
			u := UserInfo{
				username: username,
				password: password,
			}
			fmt.Printf("%#v\n", u)
		*/
		var u UserInfo
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"msg": "ok",
			})
		}
	})

	router.POST("/user", func(c *gin.Context) {
		var u UserInfo
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"msg": "ok",
			})
		}
	})
	// bind with json
	router.POST("/json", func(c *gin.Context) {
		var u UserInfo
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"msg": "ok",
			})
		}
	})

	router.Run(":9090")
}

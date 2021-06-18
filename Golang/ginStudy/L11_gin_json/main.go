/*
 * @Author: Liu Yuchen
 * @Date: 2021-03-31 17:38:14
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-03-31 17:58:53
 * @Description: Say Something, Motherfuncker
 * @FilePath: /go/src/goStudy/gin_demo/L11_gin_json/main.go
 * @GitHub: https://github.com/liuyuchen777
 */

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLFiles("./templates/index.tmpl")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", nil)
	})

	router.GET("/json", func(c *gin.Context) {
		// method 1: map
		data := gin.H{
			"name":   "lyc",
			"age":    22,
			"gender": "male",
		}
		c.JSON(http.StatusOK, data)
	})

	type stu struct {
		// 灵活使用tag进行定制化操作
		// 如果需要属性小写，打tag
		// golang反射机制
		Name   string `json:"name"`
		Age    int    `json:"age"`
		Gender string `json:"gender"`
	}

	router.GET("/struct", func(c *gin.Context) {
		// method 2: struct
		data := stu{
			// 必须首字母大写
			Name:   "lyc",
			Age:    22,
			Gender: "male",
		}
		c.JSON(http.StatusOK, data)
	})

	router.Run(":9090")
}

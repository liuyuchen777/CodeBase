/*
 * @Author: Liu Yuchen
 * @Date: 2021-04-01 10:09:21
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-04-02 00:56:53
 * @Description: Say Something, Motherfuncker
 * @FilePath: /go/src/goStudy/ginStudy/L16_file_upload/main.go
 * @GitHub: https://github.com/liuyuchen777
 */

package main

import (
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLFiles("./index.html")

	router.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.POST("/upload", func(c *gin.Context) {
		// 从请求中读取文件
		f, err := c.FormFile("f1") // 从请求中获取携带的参数
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": err.Error(),
			})
		} else {
			// 姜读取到的文件保存在本地
			// dst := fmt.Sprintf("./%s", f.Filename)
			dst := path.Join("./", f.Filename)
			c.SaveUploadedFile(f, dst)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})

	router.Run(":9090")
}

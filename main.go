package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func writLog(name string) {
	file, err := os.OpenFile("./log.text", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件出错，err:", err)
		return
	}
	fmt.Fprintln(file, name)
}

func main() {
	// 创建一个路由
	r := gin.Default()

	// 设置 HTML 模板路径
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	api := r.Group("/api")
	{
		api.GET("/upload/log", func(c *gin.Context) {
			currentTime := time.Now().Format("2006-01-02 15:04:05")

			fmt.Println(currentTime)
			fmt.Println(currentTime)

			// 写入到本地文件
			queryString := currentTime + ":" + c.Query("v")
			// requestTime := c.Request.URL

			writLog(queryString)

			// 返回一个文本
			c.String(http.StatusOK, "done")
		})
	}

	// 启动一个http服务
	r.Run()
}

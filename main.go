package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func writLog(name string) {
	file, err := os.OpenFile("./log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件出错，err:", err)
		return
	}
	fmt.Fprintln(file, name)
}

func main() {
	// 创建一个路由
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/upload/log", func(c *gin.Context) {
			currentTime := time.Now().Format("2006-01-02 15:04:05")

			// 写入到本地文件
			queryString := "\n---------------" + currentTime + ":" + c.Query("v")

			writLog(queryString)

			// 返回一个文本
			c.String(http.StatusOK, "done")
		})

		api.GET("/upload/log/clear", func(c *gin.Context) {
			filePath := "log.txt"

			// 使用 Remove 函数删除文件
			err := os.Remove(filePath)
			if err != nil {
				fmt.Println("Error:", err)
				c.String(http.StatusOK, "删除失败")
				return
			}

			c.String(http.StatusOK, "File deleted successfully.")
		})
	}

	// 启动一个http服务
	r.Run()
}

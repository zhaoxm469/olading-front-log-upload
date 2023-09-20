package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Log struct {
	Content     string `json:"content"`
	CurrentTime string `json:"currentTime"`
}

func createOrUpdateJSONFile(filename string, data interface{}) error {

	// 尝试打开已有的 JSON 文件以读取数据
	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil && os.IsNotExist(err) {
		// 如果文件不存在，则创建文件并初始化为一个空的 JSON 数组
		file, err = os.Create(filename)
		if err != nil {
			return err
		}
		defer file.Close()

		emptyData := []interface{}{}
		encoder := json.NewEncoder(file)
		err = encoder.Encode(emptyData)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	defer file.Close()

	// 解析已有的 JSON 数据
	var existingData []interface{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&existingData)
	if err != nil {
		return createOrUpdateJSONFile(filename, data)
	}

	// 将新的数据追加到切片中
	existingData = append(existingData, data)

	// 将更新后的数据重新写入文件
	file.Seek(0, 0)
	encoder := json.NewEncoder(file)
	err = encoder.Encode(existingData)
	if err != nil {
		return err
	}

	return nil
}

func contains(s string, arr []string) bool {
	for _, a := range arr {
		if strings.Contains(s, a) {
			return true
		}
	}
	return false
}

func main() {

	// 创建一个路由
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/upload/log", func(c *gin.Context) {
			currentTime := time.Now().Format("2006-01-02 15:04:05")
			queryString := c.Query("v")
			filename := "log.json"

			newLog := Log{
				Content:     queryString,
				CurrentTime: currentTime,
			}

			exclusions := []string{"/marketing/dict/bankList", "/marketing/merchant/orgTree"}

			if contains(queryString, exclusions) {
				fmt.Println("字符串包含数组内的字符")
			} else {
				err := createOrUpdateJSONFile(filename, newLog)
				if err != nil {
					c.String(http.StatusOK, "失败")
					fmt.Println("Error:", err)
					return
				}
			}

			pngFilePath := "./bitbug_favicon.ico" // 替换成你的 PNG 图片文件路径
			pngData, err := ioutil.ReadFile(pngFilePath)
			if err != nil {
				c.String(http.StatusInternalServerError, "无法读取图片文件")
				return
			}
			// 将 PNG 图片数据进行 base64 编码
			pngBase64 := base64.StdEncoding.EncodeToString(pngData)
			// 返回 base64 编码的 PNG 图片作为响应
			c.Data(http.StatusOK, "image/png", []byte(pngBase64))
		})

		api.GET("/upload/log/clear", func(c *gin.Context) {
			filePath := "log.json"

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

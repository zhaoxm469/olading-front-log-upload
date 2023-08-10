package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Log struct {
	Content string `json:"content"`
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

func main() {
	newLog := Log{
		Content: "Alice",
	}
	filename := "log.json"

	err := createOrUpdateJSONFile(filename, newLog)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Data appended to or created in", filename)
}

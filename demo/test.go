package main

import (
	"fmt"
	"os"
)

func main() {
	// 向标准输出写入内容
	// fmt.Fprintln(os.Stdout, "向标准输出写入内容")
	// fileObj, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	// if err != nil {
	// 	fmt.Println("打开文件出错，err:", err)
	// 	return
	// }
	// // 向打开的文件句柄中写入内容
	// fmt.Fprintln(fileObj, "往文件中写如信息")
	// file, err := os.Open("./xx2.txt")
	// if err != nil {
	// 	fmt.Println("打开文件出错，err:", err)
	// }
	// file.Close()

	name := 123

	fileObj, err := os.OpenFile("./11.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件出错，err:", err)
	}

	defer fileObj.Close()

	fmt.Fprintf(fileObj, "往文件中写如信息：%v", name)
}

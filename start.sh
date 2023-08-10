#!/bin/bash

# 编译 Go 项目
go build main.go

# 使用 pm2 启动项目
pm2 start ./main --name go-olading-upload-log

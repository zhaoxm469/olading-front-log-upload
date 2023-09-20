#!/bin/bash

# 停止项目
pm2 delete go-olading-upload-log

# 编译 Go 项目
go build main.go

# 使用 pm2 启动项目
pm2 start ./main --name go-olading-upload-log

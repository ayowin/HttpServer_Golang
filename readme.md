[toc]

# HttpServer_Golang
golang基于gin的轻量级http框架

## 1. 环境说明
* 本工程使用vscode搭建
* mysql，需自行建库并导入数据（**http-server-golang.sql**)

## 2. 依赖库
* vscode相关的库需自行安装（在使用时，根据vscode的提示安装即可）
  ```golang
  示例：
  go install -v github.com/ramya-rao-a/go-outline@latest
  go install -v github.com/go-delve/cmd/dlv@latest
  ...
  ```
* gin安装：
  ```golang
  go get github.com/gin-gonic/gin
  ```
* mysql: 
  ```golang
  go get github.com/go-sql-driver/mysql
  ```

## 3. 运行、编译
* 运行
  ```golang
  go run main.go
  ```
* 编译
  ```golang
  # 编译
  go build

  # 清理
  go clean
  ```
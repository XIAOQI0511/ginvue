# go 执行顺序

`go run main.go`

`go mod init qi.xiao/ginessential`

`go get -u github.com/gin-gonic/gin`

`go get -u github.com/jinzhu/gorm`

`go get github.com/dgrijalva/jwt-go`

`go get github.com/spf13/viper`
config 组件

# 使用工具

## postman

create a request
body form-data

顺序：
1.model 里添加结构体
2.controller 创建 db，添加增删改查方法
3.routes 里添加路由
4.postman 测试 body raw 选 JSON

重构：
创建新的模型来承接参数
定义 validater

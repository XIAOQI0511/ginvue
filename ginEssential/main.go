package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"qi.xiao/ginessential/common"
)

func main() {
	db := common.InitDB()
	defer db.Close() //延迟关闭？
	r := gin.Default()
	r = CollectRoute(r)
	panic(r.Run(":1016"))
}

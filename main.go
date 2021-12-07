package main

import (
	"github.com/gin-gonic/gin"
	"qdgo/goweb/common"
)


func  main()  {
		//defer db.Close 可能最新版的gorm 不需要关闭了
		common.InitDB()
		r := gin.Default()
		r =  CollectRouter(r)
		panic(r.Run()) // 监听并在 0.0.0.0:8080 上启动服务
}


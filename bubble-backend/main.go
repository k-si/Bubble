package main

import (
	"bubble/dao"
	"bubble/model"
	"bubble/router"
)

// 结构体和sql表进行一次映射，可以实现创建和更新sql表
func migrateModels() {
	if err := dao.DB.AutoMigrate(model.Bubble{}); err != nil {
		panic("migrate models failed")
	}
}

func main() {
	// 连接数据库
	dao.ConnMySQL()

	migrateModels()

	// 启动服务
	router.Run()
}

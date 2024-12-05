package main

import (
	"attendance_uniapp/initializer"
	"attendance_uniapp/routes"
)

func main() {
	// InitConfig 初始化配置
	initializer.Init()
	r := routes.SetupRouter()
	_ = r.Run(initializer.GlobalConfig.App.Port)
}

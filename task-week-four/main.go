package main

import (
	"github.com/spf13/viper"
	"log/slog"
	"task-week-four/handlers"
	"task-week-four/models"
	"task-week-four/utils"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	utils.ParseConfig()
	utils.SetMode()
	utils.SetLogger()
	utils.InitDB() // 初始化数据库
	models.Init()  // 迁移数据库，初始化模型
	r := handlers.InitEngine()
	server_port := viper.GetString("app.addr")
	slog.Info("service  is listenning", "addr", server_port)
	println("server port: " + server_port)
	r.Run(server_port)
}

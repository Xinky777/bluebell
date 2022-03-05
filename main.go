package main

import (
	"fmt"
	"web_app/controller"
	"web_app/dao/mysql"
	"web_app/dao/redis"
	"web_app/logger"
	"web_app/pkg/snowflake"
	"web_app/router"
	"web_app/settings"

	"go.uber.org/zap"
)

//go Web开发通用的脚手架模版

func main() {
	//1.加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("init setting failed,err:%v\n", err)
		return
	}

	//2.初始化日志
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("Init logger failed,err:%v\n", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success...")

	//3.初始化MySQL连接
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("Init redis failed,err:%v\n", err)
		return
	}
	defer mysql.Close()

	//4.初始化redis连接
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("Init redis failed,err:%v\n", err)
		return
	}
	defer redis.Close()

	//雪花算法生成随机ID
	if err := snowflake.Init(settings.Conf.StartTime, settings.Conf.MachineID); err != nil {
		fmt.Printf("Init snowflake failed,err:%v\n", err)
		return
	}

	//初始化gin框架内置的翻译器
	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("init trans failed,err:%v\n", err)
		return
	}

	//5.注册路由
	r := router.Setup(settings.Conf.Mode)
	s := settings.Conf.Port
	fmt.Printf("端口号：%v\n", s)
	if err := r.Run(fmt.Sprintf(":%d", settings.Conf.Port)); err != nil {
		fmt.Printf("run server failed,err:%v\n", err)
		return
	}
}

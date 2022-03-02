package logic

import (
	"web_app/dao/mysql"
	"web_app/pkg/snowflake"
)

//存放业务逻辑的代码

func SingUp() {
	//判断用户是否存在
	mysql.QueryUserByUsername()
	//生成UID
	snowflake.GenID()
	//保存进数据去
	mysql.InsertUser()
}

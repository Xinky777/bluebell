package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/jwt"
	"bluebell/pkg/snowflake"
)

//存放业务逻辑的代码

//SignUp 注册业务逻辑
func SignUp(p *models.ParamSignUp) (err error) {
	//1.判断用户是否存在
	if err = mysql.CheckUserExist(p.Username); err != nil {
		//数据库查询出错
		return err
	}

	//2.生成UID
	userID := snowflake.GenID()
	//构造一个User实例
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	//3.保存进数据库
	return mysql.InsertUser(user)
}

//Login 登录业务逻辑
func Login(p *models.ParamLogin) (user *models.User, err error) {
	//1.将用户输入的参数存放到user结构体中 用于后续步骤与数据库中数据校验
	user = &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	//传递的是指针 能够拿到user.UserID
	if err = mysql.Login(user); err != nil {
		return nil, err
	}
	//生成jwt
	token, err := jwt.GenToken(user.UserID, user.Username)
	if err != nil {
		return
	}
	user.Token = token
	return
}

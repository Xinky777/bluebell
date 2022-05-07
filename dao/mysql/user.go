package mysql

import (
	"bluebell/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
)

//把每一步数据库操作封装成函数
//待logic层根据业务需求调用

const secret = "baidu.com"

//CheckUserExist 检查指定用户名的用户是否存在
func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err = db.Get(&count, sqlStr, username); err != nil {
		//数据库查询错误
		return err
	}
	if count > 0 {
		//用户已存在
		return ErrorUserExist
	}
	return
}

// InsertUser 像数据库中插入一条新的用户记录
func InsertUser(user *models.User) (err error) {
	//对密码进行加密
	user.Password = encryptPassword(user.Password)
	//执行SQL语句入库
	sqlStr := `insert into user(user_id,username,password) values(?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

//encryptPassword 将原始用户输入的密码加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

//Login 登陆
func Login(user *models.User) (err error) {
	oPassword := user.Password //用户输入的密码
	//1.判断数据库里是否有此人
	sqlStr := `select user_id, username, password from user where username=?`
	err = db.Get(user, sqlStr, user.Username) //从数据库中查询结果 并返回user结构体
	if err == sql.ErrNoRows {
		//查询数据库成功 用户不存在
		return ErrorUserNotExist
	}
	if err != nil {
		//查询数据库失败
		return err
	}
	//2.判断密码是否正确
	password := encryptPassword(oPassword) //加密后的密码 数据库里存储的是加密后的密码
	if password != user.Password {         //此处的user.Password是数据库中存储的密码
		return ErrorInvalidPassword
	}
	//密码验证成功 返回
	return
}

//GetUserById 根据id获取用户信息
func GetUserById(UserId int64) (User *models.User, err error) {
	User = new(models.User)
	sqlStr := `select user_id,username
				from user
				where user_id = ?`
	err = db.Get(User, sqlStr, UserId)
	return
}

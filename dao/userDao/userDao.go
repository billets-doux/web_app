package userDao

import (
	"errors"
	"web_app/dao/mysql"
	"web_app/models/user"
)

// InsertUser 向数据库插入一条新的用户记录
func InsertUser(user *user.User) (err error) {
	sqlStr := `insert into user (user_id,username,password) values(?,?,?)`
	_, err = mysql.DB.Exec(sqlStr, user.UserID, user.UserName, user.Password)
	return err
}

func CheckUserExists(username string) error {
	sqlStr := `select count(1) from user where username = ?`
	var count int
	if err := mysql.DB.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已存在")
	}
	return nil
}

func CheckUserByUsernameAndPassword(user *user.User) (err error) {
	sqlStr := `select user_id,username,password from user where username = ? and password = ?`

	if err := mysql.DB.Get(user, sqlStr, user.UserName, user.Password); err != nil {
		return err
	}
	if user.UserID == 0 {
		return errors.New("用户名或密码错误")
	}
	return nil

}

package userService

import (
	"crypto/md5"
	"encoding/hex"
	"web_app/dao/userDao"
	"web_app/models"
	"web_app/models/user"
	"web_app/utils/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	// 1.判断用户是否存在
	if err = userDao.CheckUserExists(p.UserName); err != nil {
		return err
	}
	// 2.生成UID
	userID, _ := snowflake.GenID()
	u := user.User{
		UserID:   userID,
		UserName: p.UserName,
		Password: encryptPassword(p.Password),
	}
	// 3.保存到数据库
	err = userDao.InsertUser(&u)
	return err
}

func Login(p *models.ParamLogin) (err error) {
	//1.判断用户是否存在
	if err = userDao.CheckUserExists(p.UserName); err == nil {
		return err
	}
	//2.检查用户名密码是否匹配
	u := user.User{
		UserName: p.UserName,
		Password: encryptPassword(p.Password),
	}
	err = userDao.CheckUserByUsernameAndPassword(&u)
	return err
}

// encryptPassword 密码加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	salt := "billetsdoux"
	h.Write([]byte(salt))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

package orm

import (
	"todoList/utils"
)

func GetUser(by, info string) (*User, bool) {
	var u User
	result := DB.Where(by+" = ?", info).First(&u)
	return &u, result.Error == nil
}

func CreateUser(name, email, pwd string) (u *User, ok bool, msg string) {
	hpwd, err := utils.GenHashPassword(pwd)
	if err != nil {
		return nil, false, "please register again"
	}
	u = &User{Name: name, Password: hpwd, Email: email, TodoNum: 0}
	result := DB.Create(u)
	if result.Error != nil {
		return nil, false, "username has been used"
	}
	return u, true, "ok"
}

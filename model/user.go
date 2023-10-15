package model

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       int    `json:"id"`
	UserName string `json:"userName"`
	FName    string `json:"fName"`
	SName    string `json:"sName"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
	Password string `json:"password"`
}

// CreateUser - Add the new User to database
// input - UserId and password
// err - error
func (u *User) CreateUser(userName string, password string) (user *User, err error) {
	user.UserName=userName
	user.Password=password
	o:=orm.NewOrm()
	_, err = o.Insert(&user)
	return
}

// find the user by its id
// input - userId
// return - user
// err - error
func (u *User) GetUserById(userId int) (user *User, err error) {
	o := orm.NewOrm()
	err = o.QueryTable("user").Filter("id", userId).One(&user)
	return
}

// find the user by userName
// input - userName
// return - user
// err - error
func (u *User) GetUserByUserName(userName string) (user *User, err error) {
	o := orm.NewOrm()
	err = o.QueryTable(new(User)).Filter("username", userName).One(&user)
	return
}

func (u *User) UpdateUser() (id int, err error) {
	return
}
func (u *User) DeleteUser() (id int, err error) {
	return
}

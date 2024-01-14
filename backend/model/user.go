package model

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       int64    `json:"id"`
	UserName string `json:"userName"`
	FName    string `json:"fName"`
	SName    string `json:"sName"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
	Password string `json:"password"`
}

func (u *User) TableName() string {
	return "ls_user"
}
// CreateUser - Add the new User to database
// input - UserId and password
// err - error
func (u *User) CreateUser(user User) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(&user)
	return
}

// find the user by its id
// input - userId
// return - user
// err - error
func (u *User) GetUserById(userId int) (user *User, err error) {
	o := orm.NewOrm()
	err = o.QueryTable(new(User)).Filter("id", userId).One(&user)
	return
}

// find the user by userName
// input - userName
// return - user
// err - error
func (u *User) GetUserByUserName(userName string) (user *User, err error) {
	o := orm.NewOrm()
	err = o.QueryTable(new(User)).Filter("username", userName).One(&user)
	return user, nil
}

// Update the user by its id
// and return last inserted id
// else return error 
func (u *User) UpdateUserbyId(userId int64) (id int64, err error) {
	o := orm.NewOrm()
	user := User{Id: userId}
    id, err = o.Update(&user)
	return
}
// Delete the user by its id 
// TODO call the ls_delete_user to store the delete user
func (u *User) DeleteUserbyId(userId int64) (id int64, err error) {
	o:=orm.NewOrm()
	user:=User{Id:userId}
	// call the delete user to store the user details

	id, err =o.Delete(&user)
	return
}
// add the delete user
// and return last inserted id else return err
func (u *User) AddDeleteUser(userId int64)(id int64,err error){
	o:=orm.NewOrm()
	user:=User{Id: userId}
	id, err =o.Insert(&user)
    return id, err
}
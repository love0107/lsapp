package model

import (
	"errors"
	"time"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Id        int64     `orm:"column(id);pk;auto"`
	FName     string    `orm:"column(fName);size(255)"`
	SName     string    `orm:"column(sName);size(255)"`
	Mobile    string    `orm:"column(mobile);size(20);unique"`
	Email     string    `orm:"column(email);size(255);unique"`
	Gender    string    `orm:"column(gender);size(10)"`
	Password  string    `orm:"column(password);size(255)"`
	CreatedOn time.Time `orm:"column(createdOn);type(datetime);auto_now_add"`
	UpdatedOn time.Time `orm:"column(updatedOn);type(datetime);auto_now"`
}

func (u *User) TableName() string {
	return "ls_user"
}

// CreateUser - Add the new User to database
// input - UserId and password
// err - error
func (u *User) CreateUser() (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(u)
	return
}

// find the user by its id
// find the user by its id
// input - userId
// return - user
// err - error
func (u *User) GetUserById(userId int) (user *User, err error) {
	o := orm.NewOrm()
	user = &User{} // Initialize user variable
	err = o.QueryTable(new(User)).Filter("id", userId).One(user)
	return user, err
}

// find the user by userName
// input - userName
// return - user
// err - error
func (u *User) GetUserByUserName(userName string) (user *User, err error) {
	o := orm.NewOrm()
	user = &User{}
	err = o.QueryTable(new(User)).Filter("userName", userName).One(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// find the user by email
// input - userName
// return - user
// err - error
func (u *User) GetUserByEmail(email string) (user *User, err error) {
	o := orm.NewOrm()
	user = &User{}
	err = o.QueryTable(new(User)).Filter("email", email).One(user)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("User not found")
	}
	return user, nil
}

// find the user by mobile
// input - mobile
// return - user
// err - error
func (u *User) GetUserByMobile(mobile string) (*User, error) {
	o := orm.NewOrm()
	user := &User{}
	err := o.QueryTable(new(User)).Filter("mobile", mobile).One(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// find the user by mobile
// input - mobile
// return - user
// err - error
func (u *User) UpdateUserPassword(password string) (id int64, err error) {
	o := orm.NewOrm()
	u.Password = password
	id, err = o.Update(u)
	if err != nil {
		return id, err
	}
	return id, nil
}

// Update the user by its id
// and return last inserted id
// else return error
// func (u *User) UpdateUserbyId(userId int64) (id int64, err error) {
// 	o := orm.NewOrm()
// 	user := User{Id: userId}
//     id, err = o.Update(&user)
// 	return
// }
// Delete the user by its id
// TODO call the ls_delete_user to store the delete user
// func (u *User) DeleteUserbyId(userId int64) (id int64, err error) {
// 	o:=orm.NewOrm()
// 	user:=User{Id:userId}
// 	// call the delete user to store the user details

// 	id, err =o.Delete(&user)
// 	return
// }
// add the delete user
// and return last inserted id else return err
// func (u *User) AddDeleteUser(userId int64)(id int64,err error){
// 	o:=orm.NewOrm()
// 	user:=User{Id: userId}
// 	id, err =o.Insert(&user)
//     return id, err
// }

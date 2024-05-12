package model

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// Session represents a user session in the application.
type Session struct {
	Id        int64     `orm:"column(id);pk;auto"`
	SessionID string    `orm:"column(sessionID);size(255)"`
	UserID    int64     `orm:"column(userID);"`
	CreatedOn time.Time `orm:"column(createdOn);type(datetime);auto_now_add"`
	UpdatedOn time.Time `orm:"column(updatedOn);type(datetime);auto_now"`
}

func (s *Session) TableName() string {
	return "ls_session"
}

func (s *Session) GetSessionByUserID(userID int64) (*Session, error) {
    o := orm.NewOrm()
    session := &Session{}
    err := o.QueryTable(new(Session)).Filter("userID", userID).OrderBy("-Id").Limit(1).One(session)
    if err != nil {
        return nil, err
    }
    return session, nil
}

package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type User struct {
	Id         int        `orm:"pk;auto" json:"id"`
	UserName   string     `orm:"description(用户名);index;unique" json:"user_name"`
	Password   string     `orm:"description(密码)" json:"password"`
	Sex        int        `orm:"description(性别,0是女生,1是男生);default(1)" json:"sex"`
	Age        int        `orm:"description(年龄)" json:"age"`
	IsAdmin    int        `orm:"description(1是管理员,2是普通用户);default(2)" json:"is_admin"`
	CreateTime time.Time  `orm:"auto_now_add;type(datetime);description(注册时间)" json:"create_time"`
	Cover      string     `orm:"description(头像);default(static/upload/bq3.png)" json:"cover"`
	Blogs      []*Blog    `orm:"reverse(many)" json:"blogs"`    //一对多
	Comments   []*Comment `orm:"reverse(many)" json:"comments"` //一对多
}

func (u *User) TableName() string {
	return "auth_user"
}

func init() {
	orm.RegisterModel(new(User), new(Blog), new(Comment))
}

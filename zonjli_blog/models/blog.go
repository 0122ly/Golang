package models

import (
	"time"
)

type Blog struct {
	Id         int        `orm:"pk;auto" json:"id"`
	Title      string     `orm:"description(博客标题)" json:"title"`
	Desc       string     `orm:"description(博客描述)" json:"desc"`
	Content    string     `orm:"size(4000);description(博客内容)" json:"content"`
	Cover      string     `orm:"description(博客封面);default(static/upload/no_pic.jpg" json:"cover"`
	ReadNum    int        `orm:"description(博客阅读数);default(0)" json:"read_num"`
	StarNum    int        `orm:"description(博客点赞数);default(0)" json:"star_num"`
	Author     *User      `orm:"description(博客作者);rel(fk)" json:"author"`
	Comments   []*Comment `orm:"reverse(many)" json:"comments"` //一对多
	CreateTime time.Time  `orm:"auto_now_add;type(datetime);description(创建时间)" json:"create_time"`
}

func (b *Blog) TableName() string {
	return "auth_blog"
}

package models

import (
	"time"
)

type CommentTree struct {
	Id         int
	Content    string
	Author     *User
	CreateTime time.Time
	Children   []*CommentTree
}

//评论表
type Comment struct {
	Id                 int       `orm:"pk;auto" json:"id"`
	Content            string    `orm:"size(100);description(评论内容)" json:"content"`
	Author             *User     `orm:"rel(fk);description(评论人)" json:"author"`
	Blog               *Blog     `orm:"rel(fk);description(博客外键)" json:"blog"`
	PId                int       `orm:"description(父级评论id);default(0)" json:"p_id"`
	PUserId            *User     `orm:"rel(fk);description(父级评论user_id)" json:"p_user_id"`
	ReplyCommentUserId int       `orm:"description(被回复的评论user_id)" json:"reply_comment_user_id"`
	ReplyCommentId     int       `orm:"description(被回复的评论id)" json:"reply_comment_id"`
	CommentLevel       int       `orm:"description(回复文章的都是1级评论,其他的都是二级评论)" json:"comment_level"`
	Status             int       `orm:"description(评论状态,评论被删除了状态为0);default(1)" json:"status"`
	PraiseNum          int       `orm:"description(点赞数量)" json:"praise_num"`
	TopStatus          bool      `orm:"description(评论是否置顶)" json:"top_status"`
	CreateTime         time.Time `orm:"auto_now_add;type(datetime);description(评论时间)" json:"create_time"`
}

func (c *Comment) TableName() string {
	return "auth_blog_comment"
}

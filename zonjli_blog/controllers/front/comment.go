package front

import (
	"zonjli_blog/models"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type CommentController struct {
	beego.Controller
}

func (c *CommentController) Post() {
	blog_id, err := c.GetInt("blog_id")

	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 500, "msg": "id参数错误"}
		c.ServeJSON()
		return
	}

	o := orm.NewOrm()
	blog := models.Blog{}
	o.QueryTable(new(models.Blog)).Filter("id", blog_id).One(&blog)

	content := c.GetString("content")
	user_name := c.GetSession("front_user_name")

	if user_name == nil {
		c.Data["json"] = map[string]interface{}{"code": 401, "msg": "未登录"}
		c.ServeJSON()
		return
	}

	user := models.User{}
	o.QueryTable(new(models.User)).Filter("user_name", user_name).One(&user)

	pid, err1 := c.GetInt("pid")

	if err1 != nil {
		pid = 0
	}

	comment := models.Comment{
		Blog:    &blog,
		Content: content,
		PId:     pid,
		Author:  &user,
	}
	_, err3 := o.Insert(&comment)
	if err3 != nil {
		c.Data["json"] = map[string]interface{}{"code": 500, "msg": "评论字符超出限制，请重新评论"}
		c.ServeJSON()
		return
	}

	c.Data["json"] = map[string]interface{}{"code": 200, "msg": "评论成功"}
	c.ServeJSON()
}

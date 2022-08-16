package front

import (
	"zonjli_blog/models"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type IndexController struct {
	beego.Controller
}

func (i *IndexController) Get() {
	o := orm.NewOrm()

	blogs := models.Blog{}
	o.QueryTable(new(models.Blog)).RelatedSel().All(&blogs)
	front_user_name := i.GetSession("front_user_name")

	if front_user_name != nil {
		front_user_name = ""
	}

	i.Data["username"] = front_user_name
	i.Data["blogs"] = blogs
	i.TplName = "front/index.html"
}

func (i *IndexController) BlogDetail() {
	id, _ := i.GetInt("id")

	o := orm.NewOrm()

	blog := models.Blog{}
	//获取QuerySeter对象
	qs := o.QueryTable(new(models.Blog)).Filter("id", id) //Filter查询id查询一条
	//联表查询
	qs.RelatedSel().One(&blog)

	//阅读数+1
	qs.Update(orm.Params{"read_num": blog.ReadNum + 1})

	front_user_name := i.GetSession("front_user_name")

	if front_user_name != nil {
		front_user_name = ""
	}

	comments := []models.Comment{}
	o.QueryTable(new(models.Comment)).Filter("blog_id", id).Filter("p_id", 0).RelatedSel().All(&comments)

	comment_trees := []models.CommentTree{}
	for _, comment := range comments {
		pid := comment.Id
		comment_tree := models.CommentTree{
			Id:         comment.Id,
			Content:    comment.Content,
			Author:     comment.Author,
			CreateTime: comment.CreateTime,
			Children:   []*models.CommentTree{},
		}
		GetChild(pid, &comment_tree)
		comment_trees = append(comment_trees, comment_tree)
	}
	i.Data["username"] = front_user_name
	i.Data["blog"] = blog
	i.Data["comment_trees"] = comment_trees
	i.TplName = "front/detail.html"
}

//递归
func GetChild(pid int, comment_tree *models.CommentTree) {

	o := orm.NewOrm()
	qs := o.QueryTable(new(models.Comment))
	comments := []models.Comment{}

	_, err := qs.Filter("p_id", pid).RelatedSel().All(&comments)
	if err != nil {
		return
	}

	//查询二级以下评论
	for i := 0; i < len(comments); i++ {
		pid := comments[i].Id
		child := models.CommentTree{Id: comments[i].Id, Content: comments[i].Content, Author: comments[i].Author, CreateTime: comments[i].CreateTime, Children: []*models.CommentTree{}}
		comment_tree.Children = append(comment_tree.Children, &child)
		GetChild(pid, &child)
	}
	return
}

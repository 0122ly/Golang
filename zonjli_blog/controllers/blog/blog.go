package blog

import (
	"fmt"
	"strconv"
	"time"
	"zonjli_blog/models"
	"zonjli_blog/utils"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type BlogController struct {
	beego.Controller
}

func (b *BlogController) Get() {
	o := orm.NewOrm()
	blogs := []models.Blog{}

	qs := o.QueryTable(new(models.Blog))

	qs.RelatedSel().All(&blogs)

	count, _ := qs.Count()

	current_page, err := b.GetInt("b")

	if err != nil {
		current_page = 1
	}

	page_size := 10

	total_pages := utils.GetPageNum(count, page_size)

	//前后首码
	arround_count := 4
	left_pages, right_pages, left_has_more, right_has_more := utils.Get_pagination_data(total_pages, current_page, arround_count)
	has_pre_page, has_next_page, pre_page, next_page := utils.HasNext(current_page, total_pages)

	//100,100
	b.Data["left_pages"] = left_pages
	b.Data["left_has_more"] = left_has_more
	b.Data["page"] = current_page

	b.Data["has_pre_page"] = has_pre_page
	b.Data["pre_page"] = pre_page
	b.Data["has_next_page"] = has_next_page
	b.Data["next_page"] = next_page

	b.Data["right_pages"] = right_pages
	b.Data["right_has_more"] = right_has_more

	b.Data["num_pages"] = total_pages //总页数
	b.Data["count"] = count           //总数量
	b.Data["posts"] = blogs
	b.TplName = "blog/blog-list.html"
}

func (b *BlogController) ToAdd() {
	b.TplName = "blog/blog-add.html"
}

func (b *BlogController) DoAdd() {
	title := b.GetString("title")
	desc := b.GetString("desc")
	content := b.GetString("content")

	f, h, err := b.GetFile("cover")

	var cover string
	if err != nil {
		cover = "static/upload/no_pic.jpg"
	}
	//先对err 进行判断，err != nil产生错误则先处理错误。err = nil ,这时候声明f.Close()不会出错
	defer f.Close()

	//生成时间戳，防止重名
	timeUnix := time.Now().Unix()               //int64类型
	time_str := strconv.FormatInt(timeUnix, 10) //将int64类型转为字符串

	path := "static/upload/" + time_str + h.Filename
	//保存获取到的文件
	err1 := b.SaveToFile("cover", path)

	if err1 != nil {
		cover = "static/upload/no_pic.jpg"
	}
	cover = path
	o := orm.NewOrm()

	author := b.GetSession("blog_user_name")
	user := models.User{}

	o.QueryTable(new(models.User)).Filter("user_name", author).One(&user)
	blog := models.Blog{
		Title:   title,
		Desc:    desc,
		Content: content,
		Cover:   cover,
		Author:  &user,
	}

	_, err2 := o.Insert(&blog)

	if err2 != nil {
		fmt.Println(err2)
		b.Data["json"] = map[string]interface{}{"code": 500, "msg": err2}
		b.ServeJSON()
	}

	b.Data["json"] = map[string]interface{}{"code": 200, "msg": "添加成功"}
	b.ServeJSON()
}

func (b *BlogController) BlogDelete() {

	id, err := b.GetInt("id")
	if err != nil {
		b.Ctx.WriteString("id参数错误")
	}

	o := orm.NewOrm()
	_, err2 := o.QueryTable(new(models.Blog)).Filter("id", id).Delete()

	if err2 != nil {
		fmt.Println(err2)
		b.Ctx.WriteString("删除错误")
	}

	b.Redirect(beego.URLFor("BlogController.Get"), 302)
}

func (b *BlogController) ToEdit() {

	id, err := b.GetInt("id")
	if err != nil {
		b.Ctx.WriteString("id参数错误")
	}

	o := orm.NewOrm()

	blog := models.Blog{}
	o.QueryTable(new(models.Blog)).Filter("id", id).One(&blog)
	b.Data["blog"] = blog
	b.TplName = "blog/blog-edit.html"
}

func (b *BlogController) DoEdit() {
	o := orm.NewOrm()

	id, err := b.GetInt("id")
	if err != nil {
		b.Data["json"] = map[string]interface{}{"code": 500, "msg": "id参数错误"}
	}
	qs := o.QueryTable(new(models.Blog)).Filter("id", id)

	title := b.GetString("title")
	desc := b.GetString("desc")
	content := b.GetString("content")

	f, h, err1 := b.GetFile("cover")

	if err1 != nil {
		_, err4 := qs.Update(orm.Params{
			"title":   title,
			"desc":    desc,
			"content": content,
		})
		if err4 != nil {
			b.Data["json"] = map[string]interface{}{"code": 500, "msg": "更新失败"}
		}

		b.Data["json"] = map[string]interface{}{"code": 200, "msg": "更新成功"}
		b.ServeJSON()

	}
	defer f.Close()

	//生成时间戳
	timeUnix := time.Now().Unix()               //int64类型
	time_str := strconv.FormatInt(timeUnix, 10) //将int64类型转为字符串

	path := "static/upload/" + time_str + h.Filename

	//保存获取的文件
	err2 := b.SaveToFile("cover", path)
	if err2 != nil {
		_, err5 := qs.Update(orm.Params{
			"title":   title,
			"desc":    desc,
			"content": content,
		})
		if err5 != nil {
			b.Data["json"] = map[string]interface{}{"code": 500, "msg": "更新失败"}
		}
		b.Data["json"] = map[string]interface{}{"code": 200, "msg": "更新成功"}
		b.ServeJSON()

	}
	_, err6 := qs.Update(orm.Params{
		"title":   title,
		"desc":    desc,
		"content": content,
		"cover":   path,
	})

	if err6 != nil {
		b.Data["json"] = map[string]interface{}{"code": 500, "msg": "更新失败"}
	}
	b.Data["json"] = map[string]interface{}{"code": 200, "msg": "更新成功"}
	b.ServeJSON()
}

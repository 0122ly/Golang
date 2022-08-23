package blog

import (
	"fmt"
	"zonjli_blog/models"
	"zonjli_blog/utils"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type LoginController struct {
	beego.Controller
}

type User struct {
	Username string
	Passowrd string
}

func (l *LoginController) Get() {
	fmt.Println("--------login get-------")
	l.TplName = "blog/login.html"
}

func (l *LoginController) Post() {

	username := l.GetString("username")
	password := l.GetString("password")
	m := models.Message{}
	fmt.Println("username: " + username)
	fmt.Println("password: " + password)

	md5_pwd := utils.GetMd5(password)

	o := orm.NewOrm()

	exist := o.QueryTable(new(models.User)).Filter("userName", username).Filter("password", md5_pwd).Exist()

	if exist {
		l.SetSession("blog_user_name", username)
		m.Success()
		fmt.Println("登录成功")
	} else {
		m.Fail()
		fmt.Println("登录错误")
		//l.Redirect(beego.URLFor("LoginController.Get"), 302) //重定向至登录页
	}
	l.Data["json"] = m.Data
	l.ServeJSON()
}

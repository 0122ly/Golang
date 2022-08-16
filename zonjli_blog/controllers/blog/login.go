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

func (l *LoginController) Get() {
	l.TplName = "login.html"
}

func (l *LoginController) Post() {
	username := l.GetString("username")
	password := l.GetString("password")

	md5_pwd := utils.GetMd5(password)

	o := orm.NewOrm()

	exist := o.QueryTable(new(models.User)).Filter("user_name", username).Filter("password", md5_pwd).Exist()

	if exist {
		l.SetSession("blog_user_name", username)
		fmt.Println("登录成功")
		l.Redirect(beego.URLFor("MainController.Get"), 302)
	} else {
		l.Redirect(beego.URLFor("LoginController.Get"), 302) //重定向至登录页
	}
}

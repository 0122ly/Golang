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
	l.TplName = "blog/login.html"
}

func (l *LoginController) Post() {

	username := l.GetString("username")
	password := l.GetString("password")
	message := models.Message{}
	fmt.Println("username: " + username)
	fmt.Println("password: " + password)

	if len("username") == 0 && username == "" {
		fmt.Println("用户名不能为空")
	}

	if len("password") == 0 && password == "" {
		fmt.Println("用户名不能为空")
	}

	md5_pwd := utils.GetMd5(password)

	o := orm.NewOrm()

	exist := o.QueryTable(new(models.User)).Filter("userName", username).Filter("password", md5_pwd).Exist()

	if exist {
		l.SetSession("blog_user_name", username)
		message.Success(nil, "登陆成功")
		fmt.Println("登录成功")
	} else {
		message.Fail(500, "登录错误")
		fmt.Println("登录错误")
	}
	l.Data["json"] = message
	l.ServeJSON()
}

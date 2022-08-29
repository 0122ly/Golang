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

	exist_1 := o.QueryTable(new(models.User)).Filter("userName", username).Filter("password", md5_pwd).Exist()

	if exist_1 {
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

func (l *LoginController) ExistUser() {
	username := l.GetString("username")
	message := models.Message{}

	if len("username") == 0 && username == "" {
		fmt.Println("用户名不能为空")
	}
	o := orm.NewOrm()

	exist_2 := o.QueryTable(new(models.User)).Filter("user_name", username).Exist()

	if !exist_2 {
		l.SetSession("blog_user_name", username)
		message.Fail(username, "该用户不是管理员哦！联系管理员升级权限吧")
	} else {
		message.Success(username, "尊敬的管理员，请登录")
	}
	l.Data["json"] = message
	l.ServeJSON()
}

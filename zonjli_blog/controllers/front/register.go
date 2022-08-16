package front

import (
	"zonjli_blog/models"
	"zonjli_blog/utils"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type RegisterController struct {
	beego.Controller
}

func (r *RegisterController) Get() {
	r.TplName = "front/register.html"
}

func (r *RegisterController) Post() {
	username := r.GetString("username")
	password := r.GetString("password")
	repassword := r.GetString("repassword")

	if password != repassword && (password == " " || repassword == " ") {
		r.Ctx.WriteString("两次密码不一致，请重新输入！")
	}

	md5_password := utils.GetMd5(password)

	o := orm.NewOrm()

	user := models.User{
		UserName: username,
		Password: md5_password,
		IsAdmin:  2,
		Cover:    "static/upload/bq3.png",
	}
	_, err := o.Insert(&user)

	if err != nil {
		r.Ctx.WriteString("哎呀！用户名已被使用")
	}
	r.Redirect("login", 302) //跳转
}

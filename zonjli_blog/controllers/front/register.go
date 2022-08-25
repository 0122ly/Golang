package front

import (
	"fmt"
	"regexp"
	"zonjli_blog/models"
	"zonjli_blog/utils"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/validation"
	_ "github.com/beego/beego/v2/core/validation"
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
	message := models.Message{}

	if password != repassword && (password == " " || repassword == " ") {
		r.Ctx.WriteString("两次密码不一致，请重新输入！")
	}

	md5_password := utils.GetMd5(password)

	o := orm.NewOrm()

	user := models.User{
		UserName: username,
		Password: md5_password,
		Age:      0,
		Sex:      1,
		IsAdmin:  2,
		Cover:    "static/upload/bq3.png",
	}
	_, err := o.Insert(&user)

	if err != nil {
		r.Ctx.WriteString("哎呀！用户名已被使用")
		message.Fail(500, "用户名已被使用")
	} else {
		message.Success(username, "用户名可用")
	}
	r.Data["json"] = message
	r.ServeJSON()
}

func (r *RegisterController) ExistUser() {
	username := r.GetString("username")
	message := models.Message{}
	valid := validation.Validation{}

	// if len("username") == 0 && username == "" {
	// 	message.Fail(401, "用户名不能为空")
	// } else if
	valid.Match(username, regexp.MustCompile("/^[a-zA-Z0-9-_]$/"), "username").Message("用户名只能使用字母，数字，下划线")
	valid.MaxSize(username, 10, "username").Message("最大长度为10位")
	valid.MinSize(username, 5, "username").Message("最小长度为5位")
	o := orm.NewOrm()

	exist := o.QueryTable(new(models.User)).Filter("user_name", username).Exist()
	fmt.Println(!exist)
	if !exist {
		r.SetSession("front_user_name", username)
		message.Success(username, "用户名可注册")
	} else {
		message.Fail(username, "该用户已被注册，换个名字吧")
	}
	r.Data["json"] = message
	r.ServeJSON()
}

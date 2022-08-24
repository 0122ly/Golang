package front

import (
	"fmt"
	"zonjli_blog/models"
	"zonjli_blog/utils"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type FrontLoginController struct {
	beego.Controller
}

type User struct {
	Username string
	Passowrd string
}

func (f *FrontLoginController) Get() {
	f.TplName = "front/login.html"
}

func (f *FrontLoginController) Post() {
	username := f.GetString("username")
	password := f.GetString("password")
	message := models.Message{}

	md5_pwd := utils.GetMd5(password)

	o := orm.NewOrm()

	exist_1 := o.QueryTable(new(models.User)).Filter("user_name", username).Filter("password", md5_pwd).Exist()

	if exist_1 {
		f.SetSession("front_user_name", username)
		message.Success(nil, "登录成功")
		fmt.Println("登录成功")
	} else {
		message.Fail(500, "登录错误")
	}
	f.Data["json"] = message
	f.ServeJSON()

}

func (f *FrontLoginController) ExistUser() {
	username := f.GetString("username")
	message := models.Message{}

	if len("username") == 0 && username == "" {
		fmt.Println("用户名不能为空")
	}
	o := orm.NewOrm()

	exist_2 := o.QueryTable(new(models.User)).Filter("user_name", username).Exist()

	if exist_2 {
		f.SetSession("front_user_name", username)
		message.Success(username, "用户名可用")
	} else {
		message.Fail(500, "用户名不存在")
	}
	f.Data["json"] = message
	f.ServeJSON()
}

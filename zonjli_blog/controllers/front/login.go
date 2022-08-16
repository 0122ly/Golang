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

func (f *FrontLoginController) Get() {
	f.TplName = "front/login.html"
}

func (f *FrontLoginController) Post() {
	username := f.GetString("username")
	password := f.GetString("password")

	md5_pwd := utils.GetMd5(password)

	o := orm.NewOrm()

	exist := o.QueryTable(new(models.User)).Filter("user_name", username).Filter("password", md5_pwd).Exist()

	if exist {

		f.SetSession("front_user_name", username)
		fmt.Println("登录成功")
		f.Redirect(beego.URLFor("IndexController.Get"), 302)
	} else {
		f.Redirect(beego.URLFor("FrontLoginController.Get"), 302)
	}

}

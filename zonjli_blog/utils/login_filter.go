package utils

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

func BlogLoginFilter(ctx *context.Context) {
	blog_user_name := ctx.Input.Session("blog_user_name")

	if blog_user_name == nil {
		ctx.Redirect(302, beego.URLFor("LoginController.Get"))
	}
}

func FrontLoginFilter(ctx *context.Context) {
	front_user_name := ctx.Input.Session("front_user_name")

	if front_user_name == nil {
		ctx.Redirect(302, beego.URLFor("FrontLoginController.Get"))
	}
}

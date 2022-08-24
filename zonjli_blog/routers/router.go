package routers

import (
	"zonjli_blog/controllers/blog"
	"zonjli_blog/controllers/front"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	//blog
	beego.Router("/blog/login", &blog.LoginController{})
	beego.Router("/blog/main", &blog.MainController{})
	beego.Router("/blog/welcome", &blog.MainController{})
	beego.Router("/blog/blog_list", &blog.BlogController{})
	beego.Router("/blog/blog_to_add", &blog.BlogController{}, "get:ToAdd")
	beego.Router("/blog/blog_do_add", &blog.BlogController{}, "post:DoAdd")
	beego.Router("/blog/blog_delete", &blog.BlogController{}, "get:BlogDelete")
	beego.Router("/blog/blog_to_edit", &blog.BlogController{}, "get:ToEdit")
	beego.Router("/blog/blog_do_edit", &blog.BlogController{}, "post:ToEdit")

	//front
	beego.Router("/", &front.IndexController{})
	beego.Router("/comment", &front.CommentController{})
	beego.Router("/detail", &front.IndexController{}, "get:BlogDetail")
	beego.Router("/register", &front.RegisterController{})
	beego.Router("/login", &front.FrontLoginController{})
	beego.Router("/login_user", &front.FrontLoginController{}, "get:ExistUser")

}

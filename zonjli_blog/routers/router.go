package routers

import (
	"zonjli_blog/controllers/blog"
	"zonjli_blog/controllers/front"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	//blog
	beego.Router("/blog", &blog.LoginController{})
	beego.Router("/blogUser", &blog.LoginController{}, "get:ExistUser")
	beego.Router("/blog/main/main", &blog.MainController{})
	beego.Router("/blog/main/welcome", &blog.MainController{})
	beego.Router("/blog/main/blog_list", &blog.BlogController{})
	beego.Router("/blog/main/blog_to_add", &blog.BlogController{}, "get:ToAdd")
	beego.Router("/blog/main/blog_do_add", &blog.BlogController{}, "post:DoAdd")
	beego.Router("/blog/main/blog_delete", &blog.BlogController{}, "get:BlogDelete")
	beego.Router("/blog/main/blog_to_edit", &blog.BlogController{}, "get:ToEdit")
	beego.Router("/blog/main/blog_do_edit", &blog.BlogController{}, "post:ToEdit")

	//front
	beego.Router("/", &front.IndexController{})
	beego.Router("/comment", &front.CommentController{})
	beego.Router("/detail", &front.IndexController{}, "get:BlogDetail")
	beego.Router("/register", &front.RegisterController{})
	beego.Router("/regUser", &front.RegisterController{}, "get:ExistUser")
	beego.Router("/login", &front.FrontLoginController{})
	beego.Router("/loginUser", &front.FrontLoginController{}, "get:ExistUser")

}

package blog

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "blog/index.html"
}

func (c *MainController) Welcome() {
	c.TplName = "blog/welcome.html"
}

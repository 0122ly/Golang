package main

import (
	_ "zonjli_blog/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}


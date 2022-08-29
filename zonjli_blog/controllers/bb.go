package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type BbControllers struct {
	beego.Controller
}

func (b *BbControllers) Get() {

}

func (b *BbControllers) Bb() (firstName, middleName, lastName, nickName string) {
	firstName = "May"
	middleName = "dd"
	lastName = "ss"
	nickName = "sfa"
	return
}

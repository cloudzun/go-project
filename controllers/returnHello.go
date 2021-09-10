package controllers

import (
	"github.com/astaxie/beego"
)

type HelloController struct {
	beego.Controller
}

func (c *HelloController) Get() {
	c.Data["json"] = "Hello, DevOps!"
	c.ServeJSON()
	return
}
package controllers

import (
    "github.com/astaxie/beego"
)

type HomeController struct {
    beego.Controller
}

func (this *HomeController) Get() {
    this.Ctx.WriteString("Hello beego")
}
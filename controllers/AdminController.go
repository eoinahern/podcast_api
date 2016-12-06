package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (u *UserController) getUser() {
  id String = this.Ctx.Input.Param(":id")
  this.Data["id"] = id
}

func (u *UserController) addUser() {

}

func (u *UserController) deleteUser() {

}


//register, login, update.

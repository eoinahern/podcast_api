package controllers

import (
	"github.com/astaxie/beego"
	"log"
)

type UserController struct {
	beego.Controller
}

func (u * UserController) URLMapping(){
		u.Mapping("getUser", u.getUser());
		u.Mapping("addUser", u.addUser());
		u.Mapping("deleteUser", u.deleteUser());
}


//@router /user/:id [get]
func (u *UserController) getUser() {
  id String = this.Ctx.Input.Param(":id")
  log.Printf("%s", id)
}


//@router /user    [post]
func (u *UserController) addUser() {

}


//@router /user/:id [delete]
func (u *UserController) deleteUser() {

}


//register, login, update.

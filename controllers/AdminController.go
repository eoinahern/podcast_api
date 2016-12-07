package controllers

import (
	"log"
	"podcast_api/models"

	"github.com/astaxie/beego"
)

type AdminController struct {
	beego.Controller
}

func (a *AdminController) URLMapping() {
	a.Mapping("GetUser", a.GetUser)
}

// @router /user/:id [get]
func (a *AdminController) GetUser() {

	id := a.GetString(":id")
	log.Print(id)

	admin := &models.Admin{
		Name:     "eoin",
		Email:    "eoinpahern@yahoo.co.uk",
		Password: "pass1",
	}

	/*jsonadmin, err := json.Marshal(admin)

	if err != nil {
		log.Fatal("err converting")
	}*/

	a.Data["json"] = &admin
	a.ServeJSON()

}

/*   //@router /user    [post]
func (u *UserController) addUser() {

}*/

/*//@router /user/:id [delete]
func (u *UserController) deleteUser() {

}*/

//register, login, update.

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
	a.Mapping("GetAdmin", a.GetAdmin)
}

// @router /user/:id [get]
func (a *AdminController) GetAdmin() {

	id := a.GetString(":id")
	log.Print(id)

	admin := &models.Admin{
		Name:         "eoin",
		Email:        "eoinpahern@yahoo.co.uk",
		Password:     "pass1",
		IsRegistered: false,
		Isloggedin:   false,
	}

	/*jsonadmin, err := json.Marshal(admin)

	if err != nil {
		log.Fatal("err converting")
	}*/

	a.Data["json"] = &admin
	a.ServeJSON()

}

//register user
//@router /user    [post]
func (u *AdminController) addAdmin() {

}

//delete account. more params passed including key
//@router /user/:id [delete]
func (u *AdminController) deleteAdmin() {

}

//@router /user/login [post]
func (u *AdminController) loginAdmin() {

}

//@router /user/:id [put]
func (u *AdminController) updateAdmin() {

}

//@router /user/logout [get]
func (u *AdminController) logoutAdmin() {

}

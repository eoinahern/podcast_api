package controllers

import (
	"encoding/json"
	"log"
	"podcast_api/models"

	"github.com/astaxie/beego"
)

type AdminController struct {
	beego.Controller
}

func (a *AdminController) URLMapping() {
	a.Mapping("GetAdmin", a.GetAdmin)
	a.Mapping("AddAdmin", a.AddAdmin)
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

//@router /user [post]
func (u *AdminController) AddAdmin() {

	var admin models.Admin
	json.Unmarshal(u.Ctx.Input.RequestBody, &admin)

	log.Print(admin.Email)
	log.Print(admin.Name)
	log.Print(admin.Password)

	log.Print("add user to db return 200 status code")
	u.Ctx.ResponseWriter.WriteHeader(200)
	u.Ctx.WriteString("success")
}

//delete account. more params passed including key
//@router /user/:id  [delete]
func (u *AdminController) deleteAdmin() {

	var id string = u.GetString(":id")
	if id == "1" || id == "2" {
		u.Ctx.ResponseWriter.WriteHeader(200)
		u.Ctx.WriteString("success")
		return
	}

	u.Ctx.ResponseWriter.WriteHeader(404)
	u.Ctx.WriteString("not found")

}

/*
//@router /user/login [post]
func (u *AdminController) loginAdmin() {

}

//@router /user/:id [put]
func (u *AdminController) updateAdmin() {

}

//stuff

//@router /user/logout [get]
func (u *AdminController) logoutAdmin() {

}*/

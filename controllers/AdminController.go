package controllers

import (
	"encoding/json"
	"log"
	"podcast_api/models"
	"podcast_api/helpers"
	"github.com/astaxie/beego/orm"

	"github.com/astaxie/beego"
)

 var blah orm.Ormer

type AdminController struct {
	beego.Controller
	blah  = new(helpers.Ormhelper).GetOrm()
}

func (a *AdminController) URLMapping() {
	a.Mapping("GetAdmin", a.GetAdmin)
	a.Mapping("AddAdmin", a.AddAdmin)
	a.Mapping("DeleteAdmin", a.DeleteAdmin)
	a.Mapping("LoginAdmin", a.LoginAdmin)

}

// @router /user/:id [get]
func (a *AdminController) GetAdmin() {
	id := a.GetString(":id")
	log.Print(id)

	var admin models.Admin = models.Admin{}
	admin.Name = "eoin"
	admin.Email = "eoinpahern@yahoo.co.uk"
	admin.Password = "Pass1"

	var adminresp = &models.AdminResp{
		Status: "Success",
		Admin:  admin,
	}

	a.Data["json"] = &adminresp
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
func (u *AdminController) DeleteAdmin() {

	var id string = u.GetString(":id")
	if id == "1" || id == "2" {
		u.Ctx.ResponseWriter.WriteHeader(200)
		u.Ctx.WriteString("success")
		return
	}

	u.Ctx.ResponseWriter.WriteHeader(404)
	u.Ctx.WriteString("not found")

}

//@router /user/login [post]
func (u *AdminController) LoginAdmin() {

	var admin models.Admin
	json.Unmarshal(u.Ctx.Input.RequestBody, &admin)

	var adminresp models.AdminResp

	adminresp.Admin.Email = admin.Email
	adminresp.Admin.Password = admin.Password

	if admin.Password == "shazam" {
		adminresp.Status = "Success"
		adminresp.Admin.Isloggedin = true
	} else {
		adminresp.Status = "failed"
		adminresp.Admin.Isloggedin = false
	}

	u.Data["json"] = &adminresp
	u.ServeJSON()

}

/*
//@router /user/:id [put]
func (u *AdminController) updateAdmin() {

}

//stuff

//@router /user/logout [get]
func (u *AdminController) logoutAdmin() {

}*/

package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["podcast_api/controllers:AdminController"] = append(beego.GlobalControllerRouter["podcast_api/controllers:AdminController"],
		beego.ControllerComments{
			Method: "GetAdmin",
			Router: `/user/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["podcast_api/controllers:AdminController"] = append(beego.GlobalControllerRouter["podcast_api/controllers:AdminController"],
		beego.ControllerComments{
			Method: "AddAdmin",
			Router: `/user`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

}

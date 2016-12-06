package routers

import (
	"podcast_api/controllers"

	"github.com/astaxie/beego"
)

//create a get, put, post and delete!!!

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/User/:id", &controllers.UserController{})

}

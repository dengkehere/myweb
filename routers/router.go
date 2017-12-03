package routers

import (
	"myweb/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    //beego.Include(&controllers.TestController{})
    beego.Router("/:id:int/:name:string",&controllers.TestController{},"get:Get_Test")
	beego.Router("/:id:int/:name:string",&controllers.TestController{},"post:Post_Test")
}

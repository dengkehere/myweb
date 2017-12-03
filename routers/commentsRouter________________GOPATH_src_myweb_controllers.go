package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["myweb/controllers:TestController"] = append(beego.GlobalControllerRouter["myweb/controllers:TestController"],
		beego.ControllerComments{
			Method: "Get_Test",
			Router: `/:id:int`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}

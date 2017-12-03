package main

import (
	_ "myweb/routers"
	"github.com/astaxie/beego"
	"fmt"
)

func main() {
	beego.AppConfig.Set("dev::test","test")
	beego.AppConfig.Set("prod::cdk","dffdddf")
	fmt.Printf("dev::cdk=%s\n",beego.AppConfig.String("dev::test"))
	value ,_ := beego.GetConfig("String","dev::cdk","nil")
	fmt.Printf("dev::cdk=%s\n",value)
	beego.Run()
}


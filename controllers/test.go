package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type TestController struct {
	beego.Controller
}

// @router /:id:int [get]
//  c.Ctx.Input.Param(":id")
func (c *TestController) Get_Test() {
	fmt.Printf("pnames:%+v\n",c.Ctx.Input.Params())
	fmt.Printf("id--:%v\n",c.GetString(":id"))
	fmt.Printf("c.Input():%+v\n",c.Input())
	fmt.Printf("")
	id := c.Ctx.Input.Param(":id")
	fmt.Printf("id:%s\n",id)
	c.Ctx.WriteString(fmt.Sprintf("id = %s\n",id))
}

func (c *TestController) Post_Test() {

	type user struct {
		Id    int         `form:"-"`
		Name  string 	  `form:"name"`
		Age   int         `form:"age"`
	}

	fmt.Printf("hi iam post\n")
	fmt.Printf("pnames:%+v\n",c.Ctx.Input.Params())
	fmt.Printf("id--:%v\n",c.GetString(":id"))
	fmt.Printf("c.Input():%+v\n",c.Input())
	fmt.Printf("c.Input.get(\"age\"):%+v\n",c.Input().Get("age"))
	fmt.Printf("c.Input[\"age\"]:%+v\n",c.Input()["age"])
	id := c.Ctx.Input.Param(":id")
	fmt.Printf("id:%s\n",id)
	fmt.Printf("input.data:%+v\n",c.Ctx.Input.Data())
	fmt.Printf("c.Ctx.Input.reqbody:%+v\n",string(c.Ctx.Input.RequestBody))
	u := user{}
	if err := c.ParseForm(&u); err != nil {
		//handle error

	}
	fmt.Printf("u:%+v\n",u)
	fmt.Printf("form:%+v\n",c.Ctx.Request.Form)
	myfile ,myheader , _:= c.GetFile("filename")
	fmt.Printf("file:%+v\n header:%+v\n",myfile,myheader)
	c.SaveToFile("filename","static/"+myheader.Filename)
	c.Ctx.WriteString(fmt.Sprintf("id = %s\n",id))
}

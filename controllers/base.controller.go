package controllers

import beego "github.com/beego/beego/v2/server/web"

type Res struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

type BaseController struct {
	beego.Controller
}

func (c *BaseController) OK(data interface{}, message string) {
	res := Res{
		Code:    0,
		Message: "Success: " + message,
		Data:    data,
	}
	c.Data["json"] = res
	err := c.ServeJSON()
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}
}

func (c *BaseController) Error(message string) {
	res := Res{
		Code:    10000,
		Message: "Error: " + message,
		Data:    nil,
	}
	c.Data["json"] = res
	err := c.ServeJSON()
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}
}

func (c *BaseController) Result(code int, data interface{}, message string) {
	res := Res{
		Code:    code,
		Message: "Error: " + message,
		Data:    data,
	}
	c.Data["json"] = res
	err := c.ServeJSON()
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}
}

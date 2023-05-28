package controllers

import beego "github.com/beego/beego/v2/server/web"

type RoleAccessController struct {
	beego.Controller
}

func (c *RoleAccessController) FindById() {
	c.Ctx.WriteString("find role access by id")
}

func (c *RoleAccessController) RoleAccessList() {
	c.Ctx.WriteString("role access list")
}

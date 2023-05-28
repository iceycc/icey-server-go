package controllers

import beego "github.com/beego/beego/v2/server/web"

type AccessController struct {
	beego.Controller
}

func (c *AccessController) CreateAccess() {
	c.Ctx.WriteString("create access")
}

func (c *AccessController) DeleteById() {
	c.Ctx.WriteString("delete access by id")
}

func (c *AccessController) UpdateById() {
	c.Ctx.WriteString("update access by id")
}

func (c *AccessController) FindById() {
	c.Ctx.WriteString("find access by id")
}

func (c *AccessController) AccessList() {
	c.Ctx.WriteString("access list")
}

func (c *AccessController) ModuleList() {
	c.Ctx.WriteString("module list")
}

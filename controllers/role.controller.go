package controllers

import beego "github.com/beego/beego/v2/server/web"

type RoleController struct {
	beego.Controller
}

func (c *RoleController) CreateRole() {
	c.Ctx.WriteString("create role")
}
func (c *RoleController) DeleteById() {
	c.Ctx.WriteString("delete role by id")
}

func (c *RoleController) UpdateById() {
	c.Ctx.WriteString("update role by id")
}

func (c *RoleController) FindById() {
	c.Ctx.WriteString("find role by id")
}

func (c *RoleController) RoleList() {
	c.Ctx.WriteString("role list")
}

package controllers

import beego "github.com/beego/beego/v2/server/web"

type AccountRoleController struct {
	beego.Controller
}

// 获取账号角色列表
func (c *AccountRoleController) AccountRoleList() {
	c.Ctx.WriteString("account role list")
}

package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"icey-go-server/controllers"
)

func RoleAccessRouter() beego.LinkNamespace {
	roleAccessNs := beego.NSNamespace("/role_access",
		beego.NSCtrlGet("/:roleId", (*controllers.RoleAccessController).FindById),
		beego.NSCtrlGet("/", (*controllers.RoleAccessController).RoleAccessList),
	)
	return roleAccessNs
}

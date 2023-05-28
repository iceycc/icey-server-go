package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"icey-go-server/controllers"
)

func RoleRouter() beego.LinkNamespace {
	roleNs := beego.NSNamespace("/role",
		beego.NSCtrlPost("/", (*controllers.RoleController).CreateRole),
		beego.NSCtrlDelete("/:id", (*controllers.RoleController).DeleteById),
		beego.NSCtrlPatch("/:id", (*controllers.RoleController).UpdateById),
		beego.NSCtrlGet("/:id", (*controllers.RoleController).FindById),
		beego.NSCtrlGet("/", (*controllers.RoleController).RoleList),
	)
	return roleNs
}

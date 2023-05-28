package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"icey-go-server/controllers"
)

func AccessRouter() beego.LinkNamespace {
	accessNs := beego.NSNamespace("/access",
		// 创建权限
		beego.NSCtrlPost("/", (*controllers.AccessController).CreateAccess),
		// 删除权限
		beego.NSCtrlDelete("/:id", (*controllers.AccessController).DeleteById),
		// 更新权限
		beego.NSCtrlPatch("/:id", (*controllers.AccessController).UpdateById),
		// 查询权限
		beego.NSCtrlGet("/:id", (*controllers.AccessController).FindById),
		// 查询权限列表
		beego.NSCtrlGet("/", (*controllers.AccessController).AccessList),
		// 查询模块列表
		beego.NSCtrlGet("/moduleList", (*controllers.AccessController).ModuleList),
	)
	return accessNs
}

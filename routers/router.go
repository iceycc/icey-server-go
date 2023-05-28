package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"icey-go-server/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	accountNs := AccountRouter()
	accessNs := AccessRouter()
	roleNs := RoleRouter()
	roleAccessNs := RoleAccessRouter()
	accountRoleNs := AccountRoleRouter()
	ns := beego.NewNamespace("/api/v1",
		accountNs,
		accessNs,
		roleNs,
		roleAccessNs,
		accountRoleNs,
	)

	beego.AddNamespace(ns)
}

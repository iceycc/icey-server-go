package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"icey-go-server/controllers"
)

func AccountRoleRouter() beego.LinkNamespace {
	accountNs := beego.NSNamespace("/account_role",
		beego.NSCtrlGet("/:accountId", (*controllers.AccountRoleController).AccountRoleList),
	)
	return accountNs
}

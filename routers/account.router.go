package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"icey-go-server/controllers"
)

func AccountRouter() beego.LinkNamespace {
	accountNs := beego.NSNamespace("/account",
		// 创建账号
		beego.NSCtrlPost("/", (*controllers.AccountController).CreateAccount),
		// 删除账号
		beego.NSCtrlDelete("/:id", (*controllers.AccountController).DeleteAccountById),
		// 修改密码
		beego.NSCtrlPatch("/modify_password", (*controllers.AccountController).ModifyPassword),
		// 重置密码
		beego.NSCtrlPatch("/reset_password/:id", (*controllers.AccountController).ResetPassword),
		// 更新账号
		beego.NSCtrlPatch("/:id", (*controllers.AccountController).UpdateById),
		// 查询账号
		beego.NSCtrlGet("/:id", (*controllers.AccountController).FindById),
		// 查询账号列表
		beego.NSCtrlGet("/", (*controllers.AccountController).AccountList),
	)
	return accountNs
}

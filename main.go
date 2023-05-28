package main

import (
	beego "github.com/beego/beego/v2/server/web"
	"icey-go-server/global"
	"icey-go-server/initialize"
	_ "icey-go-server/routers"
)

func init() {
}
func main() {
	beego.BConfig.CopyRequestBody = true
	global.BEE_ORNER = initialize.BeeOrmer()
	//models.SyncDb()
	beego.Run()
}

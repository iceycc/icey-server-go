package initialize

import "github.com/beego/beego/v2/client/orm"

func BeeOrmer() orm.Ormer {
	mysql := InitMysqlOrmer()
	return mysql
}

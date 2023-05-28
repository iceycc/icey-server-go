package initialize

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"
	"icey-go-server/models"
)

import (
	// don't forget this
	_ "github.com/go-sql-driver/mysql"
)

func Init() {
	// need to register models in init
	orm.RegisterModel(new(models.Account))
	orm.RegisterModel(new(models.Role))
	orm.RegisterModel(new(models.Access))
	orm.RegisterModel(new(models.RoleAccess))
	orm.RegisterModel(new(models.AccountRole))
	var (
		err           error
		cfg           config.Configer
		mysqlUser     string
		mysqlHost     string
		mysqlPassword string
		mysqlDb       string
		mysqlPort     string
	)
	// need to register db driver
	err = orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		print("RegisterDriver error: ", err.Error())
		return
	}
	cfg, err = config.NewConfig("ini", "./conf/my_config.ini")
	if err != nil {
		logs.Error(err)
	}
	mysqlUser, err = cfg.String("mysql_user")
	mysqlHost, err = cfg.String("mysql_host")
	mysqlPassword, err = cfg.String("mysql_password")
	mysqlDb, err = cfg.String("mysql_db")
	mysqlPort, err = cfg.String("mysql_port")
	tmpl := mysqlUser + ":" + mysqlPassword + "@tcp(" + mysqlHost + ":" + mysqlPort + ")/" + mysqlDb + "?charset=utf8"
	// need to register default database
	err = orm.RegisterDataBase("default", "mysql", tmpl)
	if err != nil {
		print("RegisterDataBase error: ", err.Error())
		return
	}
}

func SyncDb() {
	// automatically build table
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		print("RunSyncdb error: ", err.Error())
		return
	}
}

func InitMysqlOrmer() orm.Ormer {
	Init()
	o := orm.NewOrm()
	return o
}

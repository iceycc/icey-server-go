package models

type RoleAccess struct {
	Public
	RoleId   int `orm:"column(role_id)" json:"roleId"`     // 角色id
	AccessId int `orm:"column(access_id)" json:"accessId"` // 资源id
}

package models

type AccountRole struct {
	Public
	AccountId int `orm:"column(account_id)" json:"accountId"` // 账号id
	RoleId    int `orm:"column(role_id)" json:"roleId"`       // 角色id
}

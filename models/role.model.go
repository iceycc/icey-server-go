package models

type Role struct {
	Public
	Title       string `orm:"column(title)" json:"title"`             // 角色名称
	Description string `orm:"column(description)" json:"description"` // 描述
}

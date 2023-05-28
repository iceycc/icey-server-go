package models

import "time"

type Public struct {
	Id       int       `orm:"column(id);auto" json:"id"`
	IsDel    int       `orm:"column(is_del);default(0)" json:"isDel"`
	CreateAt time.Time `orm:"auto_now_add;type(datetime);column(create_at)" json:"createAt"`
	UpdateAt time.Time `orm:"auto_now;type(datetime);column(update_at)" json:"updateAt"`
}

package models

import "time"

type Account struct {
	UserName string    `orm:"column(user_name);index" json:"userName"`
	Password string    `orm:"column(password)" json:"password"`
	Email    string    `orm:"column(email)" json:"email"`
	Phone    string    `orm:"column(phone)" json:"phone"`
	Platform string    `orm:"column(platform)" json:"platform"`
	IsSuper  int       `orm:"column(is_super)" json:"isSuper"`
	Status   int       `orm:"column(status)" json:"status"`
	Id       int       `orm:"column(id);auto" json:"id"`
	IsDel    int       `orm:"column(is_del);default(0)" json:"isDel"`
	CreateAt time.Time `orm:"auto_now_add;type(datetime);column(create_at)" json:"createAt"`
	UpdateAt time.Time `orm:"auto_now;type(datetime);column(update_at)" json:"updateAt"`
}

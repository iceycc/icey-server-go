package services

import (
	"fmt"
	"icey-go-server/global"
	"icey-go-server/models"
)

type AccountListRes struct {
	Total    int64            `json:"total"`
	Page     int              `json:"page"`
	PageSize int64            `json:"pageSize"`
	List     []models.Account `json:"list"`
}

type AccountService struct{}

func (a *AccountService) CreateAccount(userName, password, email, phone string) (int64, error) {
	accountInstance := &models.Account{
		UserName: userName,
		Password: password,
		Email:    email,
		Phone:    phone,
	}
	// 判断userName是否有重逢的
	var account models.Account
	err := global.BEE_ORNER.QueryTable("account").Filter("user_name", userName).One(&account)
	if err == nil {
		return 0, fmt.Errorf("user_name is exist")
	}
	var insert int64
	insert, err = global.BEE_ORNER.Insert(accountInstance)
	return insert, err
}

func (a *AccountService) AccountList(page, pageSize int) (*AccountListRes, error) {
	var accounts []models.Account
	var (
		err   error
		total int64
		sum   int64
	)
	// 查询所有的总数
	total, err = global.BEE_ORNER.QueryTable("account").Count()
	if err != nil {
		return nil, err
	}
	// 分页查询
	sum, err = global.BEE_ORNER.QueryTable("account").Limit(pageSize, (page-1)*pageSize).All(&accounts)
	if err != nil {
		return nil, err
	}
	return &AccountListRes{
		Total:    total,
		Page:     page,
		List:     accounts,
		PageSize: sum,
	}, nil
}

func (a *AccountService) FindById(id int) (interface{}, error) {
	account := &models.Account{
		Id: id,
	}
	err := global.BEE_ORNER.Read(account)
	return account, err
}

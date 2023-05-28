package controllers

import (
	"icey-go-server/services"
	"strconv"
)

type AccountController struct {
	BaseController
	accountService *services.AccountService
}

// FindById /**
// @Description find account by id
// @Param id path int true "id"
func (c *AccountController) FindById() {
	id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	var byId interface{}
	byId, err = c.accountService.FindById(id)
	if err != nil {
		c.Error(err.Error())
		return
	}
	c.OK(byId, "find account by id success")
}

// AccountList /**
// @Description account list
// @Param page formData int true "page"
// @Param pageSize formData int true "pageSize"
func (c *AccountController) AccountList() {
	page, _ := strconv.Atoi(c.GetString("page"))
	pageSize, _ := strconv.Atoi(c.GetString("pageSize"))
	list, err := c.accountService.AccountList(page, pageSize)
	if err != nil {
		c.Error(err.Error())
		return
	}
	c.OK(list, "find account list success")
}

// CreateAccount /**
// @Description create account
// @Param userName formData string true "userName"
// @Param password formData string true "password"
// @Param email formData string true "email"
// @Param phone formData string true "phone"
func (c *AccountController) CreateAccount() {
	userName := c.GetString("userName")
	password := c.GetString("password")
	email := c.GetString("email")
	phone := c.GetString("phone")
	id, err := c.accountService.CreateAccount(userName, password, email, phone)
	if err != nil {
		c.Error(err.Error())
		return
	}
	c.OK(id, "create account success")
}

func (c *AccountController) UpdateById() {
	c.Ctx.WriteString("update account by id")
}

func (c *AccountController) DeleteAccountById() {
	c.Ctx.WriteString("delete account by id")
}

func (c *AccountController) ModifyPassword() {
	c.Ctx.WriteString("modify password")
}

func (c *AccountController) ResetPassword() {
	c.Ctx.WriteString("reset password")
}

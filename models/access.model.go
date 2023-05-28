package models

type Access struct {
	Public
	ModuleName  string `orm:"column(module_name)" json:"moduleName"`  // 模块名称
	ActionName  string `orm:"column(action_name)" json:"actionName"`  // 操作名称
	Icon        string `orm:"column(icon)" json:"icon"`               // 图标
	Url         string `orm:"column(url)" json:"url"`                 // url地址
	ModuleId    int    `orm:"column(module_id)" json:"moduleId"`      // 父模块
	Sort        int    `orm:"column(sort)" json:"sort"`               // 排序
	Description string `orm:"column(description)" json:"description"` // 描述
}

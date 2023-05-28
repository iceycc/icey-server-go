
## 框架
beego https://github.com/beego/beego

## TODOS
- [ ] 项目结构：基于beego完成项目结构搭建
- [ ] 数据库：基于beego完成数据库的操作

## 数据库实体类设计
- [ ] 用户表 account.model.go
  - username 用户名
  - password 密码
  - email 邮箱
  - phone 手机号
  - platform 平台
  - is_super 是否超级管理员
  - status 状态

- [ ] 权限表 access.model.go
  - moduleName 模块名称
  - actionName 操作名称
  - url 路由
  - icon 图标
  - moduleId 父模块ID
  - sort 排序
  - description 描述
- [ ] 角色表 role.model.go
- [ ] 角色权限表 role_access.model.go
- [ ] 用户角色表 account_role.model.go



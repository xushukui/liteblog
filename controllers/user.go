package controllers

import (
	"liteblog/models"
	"liteblog/syserror"
	"strings"
)

//UserController 管理用户
type UserController struct {
	BaseController
}

// @router /reg [post]
func (c *UserController) Register() {
	name := c.GetMustString("name", "昵称不能为空!")
	email := c.GetMustString("email", "邮箱不能为空!")
	password := c.GetMustString("password", "密码不能为空!")
	password2 := c.GetMustString("password2", "密码不能为空!")
	if strings.Compare(password, password2) == 1 {
		c.Abort500(syserror.New("两次密码输入不一致", nil))
	}
	if QueryName, err := c.Dao.QueryUserByName(name); err == nil && QueryName.ID > 0 {
		c.Abort500(syserror.New("用户昵称已存在!", err))
	}
	if QueryEmail, err := c.Dao.QueryUserByName(email); err == nil && QueryEmail.ID > 0 {
		c.Abort500(syserror.New("用户邮箱已存在!", err))
	}
	if err := models.SaveUser(&models.User{
		Name:   name,
		Email:  email,
		Pwd:    password,
		Avatar: "/static/images/info-img.png",
		Role:   1,
	}); err != nil {
		c.Abort500(syserror.New("注册失败", err))
	}
	c.JSONOk("注册成功", "/user")
}

// @router /login [post]
func (c *UserController) Login() {
	email := c.GetMustString("email", "邮箱不能为空!")
	pwd := c.GetMustString("password", "邮箱不能为空!")
	user, err := c.Dao.QueryUserByEmailAndPassword(email, pwd)
	if err != nil {
		c.Abort500(syserror.New("邮箱或密码错误", err))
	}
	c.SetSession(SESSION_USER_KEY, user)
	c.JSONOk("登录成功", "/")
	// c.ServeJSON()
}

// @router /logout [get]
func (c *UserController) Logout() {
	c.MustLogin()
	c.DelSession(SESSION_USER_KEY)
	c.Redirect("/", 302)
}

// @router /setting/editor [post]
func (c *UserController) Editor() {
	editor := c.GetMustString("editor", "default")
	if !strings.EqualFold(editor, "markdown") {
		editor = "default"
	}

	if err := c.Dao.UpdateUserEditor(editor); err != nil {
		c.Abort500(err)
		return
	}
	c.User.Editor = editor
	c.SetSession(SESSION_USER_KEY, c.User)
	c.JSONOkH("更新成功", H{
		"editor": editor,
	})
}

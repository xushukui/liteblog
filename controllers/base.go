package controllers

import (
	"errors"
	"liteblog/models"
	"liteblog/syserror"

	"github.com/astaxie/beego"
	uuid "github.com/satori/go.uuid"
)

//SESSION_USER_KEY 用来保存用户session
const SESSION_USER_KEY = "liteblog"

type NestPrepare interface {
	NestPrepare()
}

//BaseController ..
type BaseController struct {
	beego.Controller
	User    models.User
	IsLogin bool
	Dao     *models.DB
}

//Prepare ..
func (c *BaseController) Prepare() {
	c.Data["Path"] = c.Ctx.Request.RequestURI
	c.Dao = models.NewDB()
	//验证用户是否登录
	c.IsLogin = false
	user, ok := c.GetSession(SESSION_USER_KEY).(models.User)
	if ok {
		c.User = user
		c.Data["User"] = user
		c.IsLogin = true
	}
	c.Data["IsLogin"] = c.IsLogin
	//判断子controller是否实现接口 NestPrepare, 如果实现就调用
	if app, ok := c.AppController.(NestPrepare); ok {
		app.NestPrepare()
	}
}

//Abort500 ..
func (c *BaseController) Abort500(err error) {
	c.Data["error"] = err
	c.Abort("500")
}

//GetMustString 从客户端接收post请求参数
func (c *BaseController) GetMustString(key, msg string) string {
	result := c.GetString(key)
	if len(result) == 0 {
		c.Abort500(errors.New(msg))
	}
	return result
}

//MustLogin 未登录要先登录
func (c *BaseController) MustLogin() {
	if !c.IsLogin {
		c.Abort500(syserror.NoLogin{})
	}
}

type H map[string]interface{}
type ResultJsonValue struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Action string      `json:"action,omitempty"`
	Count  int         `json:"count,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

//JSONOk ..
func (c *BaseController) JSONOk(msg string, actions ...string) {
	var action string
	if len(actions) > 0 {
		action = actions[0]
	}
	c.Data["json"] = &ResultJsonValue{
		Code:   0,
		Msg:    msg,
		Action: action,
	}
	c.ServeJSON()
}

//JSONOkH ..
func (c *BaseController) JSONOkH(msg string, maps H) {
	if maps == nil {
		maps = H{}
	}
	maps["code"] = 0
	maps["msg"] = msg
	c.Data["json"] = maps
	c.ServeJSON()
}

func (c *BaseController) JSONOkData(count int, data interface{}) {
	c.Data["json"] = &ResultJsonValue{
		Code:  0,
		Count: count,
		Msg:   "成功!",
		Data:  data,
	}
	c.ServeJSON()
}

//UUID uuid是Universally Unique Identifier的缩写，即通用唯一识别码。
func (c *BaseController) UUID() string {
	u, err := uuid.NewV4()
	if err != nil {
		c.Abort500(syserror.New("系统错误", err))
	}
	return u.String()
}

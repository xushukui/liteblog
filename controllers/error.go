package controllers

import (
	"liteblog/syserror"

	"github.com/astaxie/beego/logs"
)

//ErrorController ..
type ErrorController struct {
	BaseController
}

//Error404 ..
//ajax请求: json{code:, msg:, reason:error}
func (c *ErrorController) Error404() {
	c.TplName = "error/404.html"
	if c.IsAjax() {
		c.jsonerror(syserror.Error404{})
	} else {
		c.Data["content"] = "非法访问"
	}
}

//Error500 ..
func (c *ErrorController) Error500() {
	c.TplName = "error/500.html"
	// 获取c.Data["error"] 错误，默认为 UnKnowError
	var derr error
	err, ok := c.Data["error"].(error)
	if ok {
		derr = err
	} else {
		derr = syserror.UnKnowError{}
	}
	// 将error转成 syserror.Error ，转不了就默认 UnKnowError，
	var dserr syserror.Error
	if serr, ok := derr.(syserror.Error); ok {
		dserr = serr
	} else {
		dserr = syserror.New(err.Error(), nil)
	}
	//打印日志
	if dserr.ReasonError() != nil {
		logs.Info(dserr.Error(), err)
	}
	//输出
	if c.IsAjax() {
		c.jsonerror(dserr)
	} else {
		c.Data["content"] = dserr.Error()
	}
}

func (c *ErrorController) jsonerror(serr syserror.Error) {
	c.Ctx.Output.Status = 200
	c.Data["json"] = map[string]interface{}{
		"code": serr.Code(),
		"msg":  serr.Error(),
	}
	c.ServeJSON()
}

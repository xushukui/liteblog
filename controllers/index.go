package controllers

import (
	"errors"
	"liteblog/syserror"
	"time"
)

type IndexController struct {
	//结构体继承，每次执行本结构体方法之前会先调用父级结构体方法
	//类似于类的继承
	BaseController
}

// @router /appinfo [get]
func (c *IndexController) Info() {
	dbtime := c.Dao.GetDBTime()
	var dbTs = ""
	if dbtime != nil {
		dbTs = dbtime.Format("2006-01-02 15:04:05")
	}
	c.JSONOkH("ok", H{
		"app_time": time.Now().Format("2006-01-02 15:04:05"),
		"db_time":  dbTs,
	})
}

// @router / [get]
func (c *IndexController) Get() {
	limit := 10
	//如果查询到"page"值，则page是请求传来的值, 如果没有则默认为1
	page, err := c.GetInt("page", 1)
	if err != nil || page < 1 {
		page = 1
	}
	title := c.GetString("title", "")
	if c.Dao == nil {
		c.Abort500(errors.New("数据库初始化失败! "))
	}
	ns, err := c.Dao.QueryNotesByPage(page, limit, title)
	if err != nil {
		c.Abort500(err)
	}
	if ns != nil {
		c.Data["notes"] = ns
	}
	var totpage int
	totcnt, _ := c.Dao.QueryNotesCount(title)
	if totcnt%limit == 0 {
		totpage = totcnt / limit
	} else {
		totpage = totcnt/limit + 1
	}
	c.Data["totpage"] = totpage
	c.Data["page"] = page
	c.Data["title"] = title
	// fmt.Println("~~~~~~", c.Data)
	//未登录时c.User.ID是0，普通用户登录后c.User.Role是1, 管理员登录后c.User.Role是0.
	if c.User.Role == 1 || c.User.ID == 0 {
		c.TplName = "index2.html"
	} else {
		c.TplName = "index.html"
	}
}

// @router /details/:key [get]
func (c *IndexController) GetDetail() {
	key := c.Ctx.Input.Param(":key")
	note, err := c.Dao.QueryNoteByKey(key)
	if err != nil {
		c.Abort500(syserror.New("文章不存在", err))
	}
	go c.Dao.AllVisitCount(key)
	c.Data["praise"] = false
	messages, _ := c.Dao.QueryMessageForNote(note.Key)
	c.Data["messages"] = messages
	c.Data["note"] = note
	c.TplName = "details.html"
}

// @router /comment/:key [get]
func (c *IndexController) GetComment() {
	key := c.Ctx.Input.Param(":key")
	note, err := c.Dao.QueryNoteByKey(key)
	if err != nil {
		c.Abort500(syserror.New("文章不存在", err))
	}
	c.Data["note"] = note
	c.TplName = "comment.html"
}

// @router /setting [get]
func (c *IndexController) GetSetting() {
	c.TplName = "setting.html"
}

// @router /message [get]
func (c *IndexController) GetMessage() {
	messages, err := c.Dao.QueryMessageForNote("")
	if err != nil {
		c.Abort500(err)
	}
	c.Data["messages"] = messages
	c.TplName = "message.html"
}

// @router /user [get]
func (c *IndexController) GetUser() {
	c.TplName = "user.html"
}

// @router /reg [get]
func (c *IndexController) Register() {
	c.TplName = "reg.html"
}

// @router /about [get]
func (c *IndexController) About() {
	c.TplName = "about.html"
}

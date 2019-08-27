package controllers

import (
	"liteblog/models"
	"liteblog/syserror"
)

type MessageController struct {
	BaseController
}

func (c *MessageController) NestPrepare() {

}

// @router /new/?:key [post]
func (c *MessageController) NewMessage() {
	c.MustLogin()
	key := c.UUID()
	content := c.GetMustString("content", "内容不能为空")
	notekey := c.Ctx.Input.Param(":key")
	m := &models.Message{
		UserID:  int(c.User.ID),
		User:    c.User,
		Key:     key,
		NoteKey: notekey,
		Content: content,
	}
	if err := c.Dao.SaveMessage(m); err != nil {
		c.Abort500(syserror.New("保存失败!", err))
	}
	c.JSONOkH("保存成功! ", H{
		"data": m,
	})
}

// @router /count [get]
func (c *MessageController) Count() {
	count, err := c.Dao.QueryMessageForNoteCount("")
	if err != nil {
		c.Abort500(syserror.New("查询失败", err))
	}
	c.JSONOkH("查询成功! ", H{
		"count": count,
	})
}

// @router /query [get]
func (c *MessageController) Query() {
	pageno, err := c.GetInt("pageno", 1)
	if err != nil || pageno < 1 {
		pageno = 1
	}
	limit, err := c.GetInt("limit", 10)
	if err != nil || limit < 5 {
		limit = 10
	}
	datas, err := c.Dao.QueryMessageForNoteByPage("", pageno, limit)
	if err != nil {
		c.Abort500(syserror.New("查询失败", err))
	}
	c.JSONOkH("查询成功 !", H{
		"data": datas,
	})
}

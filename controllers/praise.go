package controllers

import (
	"liteblog/models"
	"liteblog/syserror"
)

type PraiseController struct {
	BaseController
}

func (c *PraiseController) NestPrepare() {
	c.MustLogin()
}

// @router /:type/:key [post]
func (c *PraiseController) Parse() {
	key := c.Ctx.Input.Param(":key")
	ttype := c.Ctx.Input.Param(":type")
	var (
		praise  int = 0
		user_id int = int(c.User.ID)
		err     error
	)
	c.Dao.Begin()
	switch ttype {
	case "message":
		var message models.Message
		if message, err = c.Dao.QueryMessageByKey(key); err != nil {
			c.Dao.Rollback()
			c.Abort500(syserror.New("点赞失败", err))
		}
		message.Praise = message.Praise + 1
		if err := c.Dao.UpdateMessage4Praise(&message); err != nil {
			c.Dao.Rollback()
			c.Abort500(syserror.New("点赞失败", err))
		}
		praise = message.Praise
	case "note":
		var note models.Note
		if note, err = c.Dao.QueryNoteByKey(key); err != nil {
			c.Dao.Rollback()
			c.Abort500(syserror.New("点赞失败", err))
		}
		note.Praise = note.Praise + 1
		if err := c.Dao.UpdateNote4Praise(&note); err != nil {
			c.Dao.Rollback()
			c.Abort500(syserror.New("点赞失败", err))
		}
		praise = note.Praise
	default:
		c.Dao.Rollback()
		c.Abort500(syserror.New("未知类型", err))
	}
	p := models.PraiseLog{
		Key:    key,
		UserID: user_id,
		Type:   ttype,
	}
	var pp models.PraiseLog
	if pp, err = c.Dao.QueryPraiseLog(key, user_id, ttype); err != nil {
		pp = p
	} else {
		if pp.Flag {
			c.Dao.Rollback()
			c.Abort500(syserror.HasPraiseError{})
		}
	}
	pp.Flag = true
	if err := c.Dao.SavePraiseLog(&pp); err != nil {
		c.Dao.Rollback()
		c.Abort500(syserror.New("点赞失败", err))
	}
	c.Dao.Commit()
	c.JSONOkH("点赞成功! ", H{
		"praise": praise,
	})
}

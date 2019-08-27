package controllers

import (
	"bytes"
	"liteblog/models"
	"liteblog/syserror"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"

	"github.com/jinzhu/gorm"
)

//NoteController 上传文章
type NoteController struct {
	BaseController
}

func (c *NoteController) NestPrePare() {
	//判断是否登录
	c.MustLogin()
	if c.User.Role != 0 {
		c.Abort500(syserror.New("您没有权限修改文章", nil))
	}
}

// @router /new [get]
func (c *NoteController) NewPage() {
	c.Data["key"] = c.UUID()
	if strings.EqualFold(c.User.Editor, "markdown") {
		c.TplName = "note_new2.html"
		return
	}
	c.TplName = "note_new.html"
}

// @router /edit/:key [get]
func (c *NoteController) EditPage() {
	key := c.Ctx.Input.Param(":key")
	note, err := c.Dao.QueryNoteByKeyAndUserId(key, int(c.User.ID))
	if err != nil {
		c.Abort500(syserror.New("文章不存在", err))
	}
	c.Data["note"] = note
	c.Data["key"] = key
	if strings.EqualFold(note.Editor, "markdown") {
		c.TplName = "note_new2.html"
		return
	}
	c.TplName = "note_new.html"
}

// @router /del/:key [post]
func (c *NoteController) Del() {
	key := c.Ctx.Input.Param(":key")
	if err := c.Dao.DelNoteByKey(key, int(c.User.ID)); err != nil {
		c.Abort500(syserror.New("删除失败", err))
	}
	c.JSONOk("删除成功!", "/")
}

// @router /save/:key [post]
func (c *NoteController) Save() {
	key := c.Ctx.Input.Param(":key")
	editor := c.GetString("editor", "default")
	title := c.GetMustString("title", "标题不能为空!")
	content := c.GetMustString("content", "内容不能为空!")
	files := c.GetString("files", "")
	summary, _ := getSummary(content)
	note, err := c.Dao.QueryNoteByKeyAndUserId(key, int(c.User.ID))
	var n models.Note
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			c.Abort500(syserror.New("保存失败!", err))
		}
		n = models.Note{
			Key:     key,
			Summary: summary,
			Title:   title,
			Files:   files,
			Content: content,
			UserID:  int(c.User.ID),
		}
	} else {
		n = note
		n.Title = title
		n.Content = content
		n.Summary = summary
		n.Files = files
		n.UpdatedAt = time.Now()
	}
	n.Editor = editor
	if strings.EqualFold(editor, "markdown") {
		n.Source = c.GetMustString("source", "内容不能为空!")
	}
	if err := c.Dao.SaveNote(&n); err != nil {
		c.Abort500(syserror.New("保存失败!", err))
	}
	c.JSONOk("保存成功", "/details/"+key)
}

func getSummary(content string) (string, error) {
	var buf bytes.Buffer
	buf.Write([]byte(content))
	doc, err := goquery.NewDocumentFromReader(&buf)
	if err != nil {
		return "", err
	}
	str := doc.Find("body").Text()
	strRune := []rune(str)
	if len(strRune) > 400 {
		strRune = strRune[:400]
	}
	return string(strRune) + "...", nil
}

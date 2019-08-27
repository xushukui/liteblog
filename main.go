package main

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"liteblog/models"
	_ "liteblog/models"
	_ "liteblog/routers"
	"os"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {
	initLog()
	initSession()
	initTemplate()
	beego.Run()
}

func initLog() {
	if err := os.MkdirAll("data/logs", 0777); err != nil {
		beego.Error(err)
		return
	}
	logs.SetLogger("file", `{"filename":"data/logs/liteblog.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)
	logs.Async(1e3)
}

//开启session
func initSession() {
	gob.Register(models.User{})
	//设置sysn请求的静态文件路径，不设置会出错(no match, 非法访问)
	beego.SetStaticPath("assert", "assert")
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "liteblog"
	//使用file保存session
	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	//设置保存session的文件路径
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "data/session"
}

func initTemplate() {
	//为名为equal的变量添加匿名函数，在模板文件里调用equal就相当于调用了此匿名函数
	beego.AddFuncMap("equrl", func(x, y string) bool {
		x1 := strings.Trim(x, "/")
		y1 := strings.Trim(y, "/")
		return strings.Compare(x1, y1) == 0
	})
	beego.AddFuncMap("eq2", func(x, y interface{}) bool {
		s1 := fmt.Sprintf("%v", x)
		s2 := fmt.Sprintf("%v", y)
		return strings.Compare(s1, s2) == 0
	})
	beego.AddFuncMap("add", func(x, y int) int {
		return x + y
	})
	beego.AddFuncMap("json", func(obj interface{}) string {
		bs, err := json.Marshal(obj)
		if err != nil {
			return "{id:0}"
		}
		return string(bs)
	})
}

package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["liteblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["liteblog/controllers:IndexController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liteblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["liteblog/controllers:IndexController"],
        beego.ControllerComments{
            Method: "About",
            Router: `/about`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liteblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["liteblog/controllers:IndexController"],
        beego.ControllerComments{
            Method: "Info",
            Router: `/appinfo`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liteblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["liteblog/controllers:IndexController"],
        beego.ControllerComments{
            Method: "GetComment",
            Router: `/comment/:key`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liteblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["liteblog/controllers:IndexController"],
        beego.ControllerComments{
            Method: "GetDetail",
            Router: `/details/:key`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liteblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["liteblog/controllers:IndexController"],
        beego.ControllerComments{
            Method: "GetMessage",
            Router: `/message`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liteblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["liteblog/controllers:IndexController"],
        beego.ControllerComments{
            Method: "Register",
            Router: `/reg`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liteblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["liteblog/controllers:IndexController"],
        beego.ControllerComments{
            Method: "GetSetting",
            Router: `/setting`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liteblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["liteblog/controllers:IndexController"],
        beego.ControllerComments{
            Method: "GetUser",
            Router: `/user`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liteblog/controllers:MessageController"] = append(beego.GlobalControllerRouter["liteblog/controllers:MessageController"],
        beego.ControllerComments{
            Method: "Count",
            Router: `/count`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liteblog/controllers:MessageController"] = append(beego.GlobalControllerRouter["liteblog/controllers:MessageController"],
        beego.ControllerComments{
            Method: "NewMessage",
            Router: `/new/?:key`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liteblog/controllers:MessageController"] = append(beego.GlobalControllerRouter["liteblog/controllers:MessageController"],
        beego.ControllerComments{
            Method: "Query",
            Router: `/query`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liteblog/controllers:NoteController"] = append(beego.GlobalControllerRouter["liteblog/controllers:NoteController"],
        beego.ControllerComments{
            Method: "Del",
            Router: `/del/:key`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liteblog/controllers:NoteController"] = append(beego.GlobalControllerRouter["liteblog/controllers:NoteController"],
        beego.ControllerComments{
            Method: "EditPage",
            Router: `/edit/:key`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liteblog/controllers:NoteController"] = append(beego.GlobalControllerRouter["liteblog/controllers:NoteController"],
        beego.ControllerComments{
            Method: "NewPage",
            Router: `/new`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liteblog/controllers:NoteController"] = append(beego.GlobalControllerRouter["liteblog/controllers:NoteController"],
        beego.ControllerComments{
            Method: "Save",
            Router: `/save/:key`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liteblog/controllers:PraiseController"] = append(beego.GlobalControllerRouter["liteblog/controllers:PraiseController"],
        beego.ControllerComments{
            Method: "Parse",
            Router: `/:type/:key`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liteblog/controllers:UploadController"] = append(beego.GlobalControllerRouter["liteblog/controllers:UploadController"],
        beego.ControllerComments{
            Method: "UploadFile",
            Router: `/uploadfile`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liteblog/controllers:UploadController"] = append(beego.GlobalControllerRouter["liteblog/controllers:UploadController"],
        beego.ControllerComments{
            Method: "UploadImg",
            Router: `/uploadimg`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liteblog/controllers:UploadController"] = append(beego.GlobalControllerRouter["liteblog/controllers:UploadController"],
        beego.ControllerComments{
            Method: "WangeditorUploadFile",
            Router: `/wangeditorfiles`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liteblog/controllers:UserController"] = append(beego.GlobalControllerRouter["liteblog/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liteblog/controllers:UserController"] = append(beego.GlobalControllerRouter["liteblog/controllers:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liteblog/controllers:UserController"] = append(beego.GlobalControllerRouter["liteblog/controllers:UserController"],
        beego.ControllerComments{
            Method: "Register",
            Router: `/reg`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liteblog/controllers:UserController"] = append(beego.GlobalControllerRouter["liteblog/controllers:UserController"],
        beego.ControllerComments{
            Method: "Editor",
            Router: `/setting/editor`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}

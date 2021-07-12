package routers

import (
	web "github.com/beego/beego/v2/server/web"
	"github.com/ms03831/todoapp-beego/controllers"
)

func init() {
	ns := 
	web.NewNamespace("/api",
		web.NSRouter("/", &controllers.MainController{}),
		web.NSRouter("/about", &controllers.AboutController{}),
		web.NSNamespace("/auth",
			web.NSRouter("/login",&controllers.UserController{},"post:Login"),
			web.NSRouter("/register",&controllers.UserController{},"post:Register"),
			web.NSRouter("/logout",&controllers.UserController{},"delete:Logout"),
			web.NSRouter("/profile",&controllers.UserController{},"get:ViewProfile"),
			web.NSRouter("/profile/create",&controllers.UserController{},"post:CreateProfile"),
			web.NSRouter("/profile/update",&controllers.UserController{},"put:UpdateProfile"),
		),
		web.NSNamespace("/todos",
			web.NSRouter("/create",&controllers.TodoController{},"put:CreateTodo"),
			web.NSRouter("/list",&controllers.TodoController{},"get:ListTodo"),
			web.NSRouter("/delete",&controllers.TodoController{},"delete:DeleteTodo"),
			web.NSRouter("/done",&controllers.TodoController{},"put:ChangeStatus"),
			web.NSRouter("/deadline",&controllers.TodoController{},"put:ChangeDeadline"),
		),
	)
	web.AddNamespace(ns)
	web.Run()
}

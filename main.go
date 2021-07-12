package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/ms03831/todoapp-beego/models"
	_ "github.com/ms03831/todoapp-beego/routers"
)

func main() {
	beego.Run()
}


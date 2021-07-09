package main

import (
	_ "hello/models"
	_ "hello/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}


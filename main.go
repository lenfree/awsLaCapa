package main

import (
	_ "github.com/lenfree/awsLaCapa/routers"

	"github.com/astaxie/beego"
)

func main() {
        beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
        if beego.BConfig.RunMode == "dev" {
                beego.BConfig.WebConfig.DirectoryIndex = true
        }
        beego.Run()
}

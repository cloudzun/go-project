package main

import (
	. "github.com/astaxie/beego"
	_ "go-project/routers"

)

func main() {
	if BConfig.RunMode == "dev" {
		BConfig.WebConfig.DirectoryIndex = true
		BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	Run()
}

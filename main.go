package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/pelletier/go-toml"
	"github.com/use-iris/configs"
	"github.com/use-iris/libs"
	"github.com/use-iris/route"
)


func main(){
	app := iris.New()
	route.Routes(app)
	err := app.Run(iris.Addr(":8080"))
	if err != nil {
		fmt.Println(err)
	}
}

func init(){
	//后面做接参数获取配置项路径
	var err error
	configs.ConfigTree,err = toml.LoadFile("github.com/use-iris/configs/config.toml")
	if err != nil{
		fmt.Println(err)
		return
	}
	libs.InitAllDB()
}

package main

import (
	"flag"
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
	err := app.Run(iris.Addr("www.local.com:8080"))
	if err != nil {
		fmt.Println(err)
	}
}

func init(){
	//后面做接参数获取配置项路径
	var configPath = flag.String("c","github.com/use-iris/configs/config.toml","defaut config")
	var err error
	configs.ConfigTree,err = toml.LoadFile(*configPath)
	if err != nil{
		fmt.Println(err)
		return
	}
	libs.InitAllDB()
}

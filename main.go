package main

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"github.com/use-iris/libs"
)
var ConfigTree *toml.Tree

func main(){

}

func init(){
	//后面做接参数获取配置项路径
	var err error
	ConfigTree,err = toml.LoadFile("github.com/use-iris/configs/config.toml")
	if err != nil{
		fmt.Println(err)
		return
	}
	InitAllDB()
}
//初始主从库连接
func InitAllDB(){
	var intconns = int(ConfigTree.Get("mysql.MaxIdleConns").(int64))
	var intOpenConns = int(ConfigTree.Get("mysql.MaxOpenConns").(int64))
	dbmaster := libs.DbConfig{
		ConfigTree.Get("mysql.master.host").(string),
		ConfigTree.Get("mysql.master.port").(string),
		ConfigTree.Get("mysql.master.database").(string),
		ConfigTree.Get("mysql.master.user").(string),
		ConfigTree.Get("mysql.master.password").(string),
		ConfigTree.Get("mysql.master.charset").(string),
		intconns,
		intOpenConns,
	}
	libs.DB_MASTER = dbmaster.InitDB()
	dbslave := libs.DbConfig{
		ConfigTree.Get("mysql.slave.host").(string),
		ConfigTree.Get("mysql.slave.port").(string),
		ConfigTree.Get("mysql.slave.database").(string),
		ConfigTree.Get("mysql.slave.user").(string),
		ConfigTree.Get("mysql.slave.password").(string),
		ConfigTree.Get("mysql.slave.charset").(string),
		intconns,
		intOpenConns,
	}
	libs.DB_SLAVE = dbslave.InitDB()
}
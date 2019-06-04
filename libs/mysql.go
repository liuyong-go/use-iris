package libs

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/use-iris/configs"
	"log"
	"os"
)

var	DbSlave *gorm.DB
var	DbMaster *gorm.DB



type DbConfig struct {
	Host         string
	Port         string
	Database     string
	User         string
	Password     string
	Charset      string
	MaxIdleConns int
	MaxOpenConns int
}

func (c *DbConfig) InitDB() *gorm.DB {

	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", c.User, c.Password, c.Host, c.Port, c.Database, c.Charset)
	db, err := gorm.Open("mysql", connString)
	if err != nil {
		log.Panic(err)
		os.Exit(-1)
	}
	db.SingularTable(true)                  //全局设置表名不可以为复数形式。
	db.DB().SetMaxIdleConns(c.MaxIdleConns) //空闲时最大的连接数
	db.DB().SetMaxOpenConns(c.MaxOpenConns) //最大的连接数
	return db
}
//初始主从库连接
func InitAllDB(){
	var intconns = int(configs.ConfigTree.Get("mysql.MaxIdleConns").(int64))
	var intOpenConns = int(configs.ConfigTree.Get("mysql.MaxOpenConns").(int64))
	dbmaster := DbConfig{
		configs.ConfigTree.Get("mysql.master.host").(string),
		configs.ConfigTree.Get("mysql.master.port").(string),
		configs.ConfigTree.Get("mysql.master.database").(string),
		configs.ConfigTree.Get("mysql.master.user").(string),
		configs.ConfigTree.Get("mysql.master.password").(string),
		configs.ConfigTree.Get("mysql.master.charset").(string),
		intconns,
		intOpenConns,
	}
	DbMaster = dbmaster.InitDB()
	dbslave := DbConfig{
		configs.ConfigTree.Get("mysql.slave.host").(string),
		configs.ConfigTree.Get("mysql.slave.port").(string),
		configs.ConfigTree.Get("mysql.slave.database").(string),
		configs.ConfigTree.Get("mysql.slave.user").(string),
		configs.ConfigTree.Get("mysql.slave.password").(string),
		configs.ConfigTree.Get("mysql.slave.charset").(string),
		intconns,
		intOpenConns,
	}
	DbSlave = dbslave.InitDB()
}

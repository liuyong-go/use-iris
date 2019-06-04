package model

import (
	"errors"
	"fmt"
	"github.com/use-iris/libs"
)

type User struct {
	//gorm.Model
	Mobile string `gorm:"type:varchar(32);not null;"`
	Nick_name string `gorm:"type:varchar(32);not null;"`
}

func (u *User) Insert(insertValue map[string]string) error{
	var uinfo User
	uinfo.Mobile = insertValue["mobile"]
	uinfo.Nick_name = insertValue["nickname"]
	if err := libs.DbMaster.Create(&uinfo).Error; err != nil {
		fmt.Print(err)
		return errors.New("新增失败")
	}
	return nil
}
func (u *User)GetList(mobile string) ([]User, error){
	var data []User
	err := libs.DbSlave.Where("mobile = ?",mobile).Find(&data).Error
	if err != nil{
		return nil,err
	}
	return data,nil
}
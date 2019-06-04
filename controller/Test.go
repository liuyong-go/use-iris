package controller

import (
	"github.com/kataras/iris"
	"github.com/use-iris/model"
)


type TestController struct {
	Cxt iris.Context
	userModel model.User
}
func (c TestController) PostTest() string{
	return "hello test"

}
func (c TestController) GetUser() string{
	return "hello user"
}
func (c TestController) Get(){
	_, _ = c.Cxt.WriteString("hello index hh")
}
func (c TestController) GetInsert() string{
	insertData  := map[string]string{"mobile":"15312377715","nickname":"liuyong"}
	err :=c.userModel.Insert(insertData)
	if err != nil{
		_, _ = c.Cxt.WriteString(err.Error())
	}
	return "insert success"
}

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
func (c TestController) GetInsert() {
	insertData := map[string]string{"mobile": "15312370015", "nickname": "aiyan"}
	err := c.userModel.Insert(insertData)
	if err != nil {
		_, _ = c.Cxt.WriteString(err.Error())
	}
	result := map[string]interface{}{
		"code":200,"status":"success",
	}
	result["data"] = insertData
	_, _ = c.Cxt.JSON(result)
}
func (c TestController) GetList(){
	mobile := c.Cxt.FormValue("mobile")
	if mobile == ""{
		c.Cxt.WriteString("手机号不能为空")
		return
	}
	data,err := c.userModel.GetList(mobile)
	if err != nil {
		c.Cxt.WriteString(err.Error())
	}
	result := map[string]interface{}{
		"code":200,
		"status":"success",
		"data":data,
	}
	c.Cxt.JSON(result)
}

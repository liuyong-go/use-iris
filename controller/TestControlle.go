package controller

import (
	"github.com/kataras/iris"
	"github.com/use-iris/configs"
	"github.com/use-iris/core"
	"github.com/use-iris/model"
)


type TestController struct {
	Rest core.RestfulController
	Cxt iris.Context
	UserModel model.User
}
func (c TestController) PostTest() string{
	return "hello test"

}
func (c TestController) GetUser(){
	cookieValue := c.Cxt.GetCookie("username",iris.CookieDecode(configs.Sc.Decode))
	_, _ = c.Cxt.WriteString(cookieValue)
}
func (c TestController) Get(){
	c.Cxt.SetCookieKV("username","ceshi",iris.CookieEncode(configs.Sc.Encode))
	_, _ = c.Cxt.WriteString("hello index hh")
}
func (c TestController) GetInsert() {
	insertData := map[string]string{"mobile": "15312370015", "nickname": "aiyan"}
	err := c.UserModel.Insert(insertData)
	if err != nil {
		_, _ = c.Cxt.WriteString(err.Error())
	}
	result := map[string]interface{}{
		"code":200,"status":"success",
	}
	result["data"] = insertData
	_, _ = c.Cxt.JSON(result)
}
func (c TestController) GetList()[]byte{
	mobile := c.Cxt.FormValue("mobile")
	if mobile == ""{
		return c.Rest.ShowError(400,"手机号不能为空")
	}
	data,err := c.UserModel.GetList(mobile)
	if err != nil {
		return c.Rest.ShowError(400,err.Error())
	}
	return c.Rest.ShowSuccess(data)
}

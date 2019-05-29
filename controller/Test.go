package controller

import "github.com/kataras/iris"

type TestController struct {
	Cxt iris.Context
}
func (c TestController) PostTest() string{
	return "hello test"

}
func (c TestController) GetUser() string{
	return "hello user"
}
func (c TestController) Get() string{
	return "hello index"
}
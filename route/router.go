package route

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/use-iris/controller"
	"github.com/use-iris/middleware"
)

func Routes(app *iris.Application){
	mvc.New(app.Party("/test", middleware.LoginAuth)).
		Handle(new(controller.TestController))
}
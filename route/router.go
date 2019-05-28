package route

import (
	"github.com/iris_demo/controller"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/use-iris/middleware"
)

func Routes(app *iris.Application){
	mvc.New(app.Party("/test", middleware.LoginAuth)).
		Handle(new(controllers.CategorysController))
}
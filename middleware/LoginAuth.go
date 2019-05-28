package middleware

import "github.com/kataras/iris/context"

func LoginAuth(Ctx context.Context){
	Ctx.Next()
}
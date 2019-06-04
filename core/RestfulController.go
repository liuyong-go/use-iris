package core

import (
	"encoding/json"
	"github.com/kataras/iris"
)

type RestfulController struct {
	Cxt iris.Context
}
func (rest RestfulController) ShowSuccess(data interface{}) []byte{
	result := map[string]interface{}{
		"code":200,
		"status":"success",
		"data":data,
	}
	rs,_ := json.Marshal(result)
	return rs
}
func (rest RestfulController) ShowError(code int64,msg string)[]byte{
	result := map[string]interface{}{
		"code":code,
		"messsage":msg,
	}
	rs,_ := json.Marshal(result)
	return rs
}

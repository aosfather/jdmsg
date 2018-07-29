package main

import (
	"github.com/aosfather/bingo/mvc"
	"fmt"
)

/**


 */

 type AlphaController struct {
 	mvc.SimpleController
 }

 func (this *AlphaController) GetUrl() string{
 	return "/jd/msg"
 }

func (this *AlphaController) GetParameType(method string) interface{} {
	return &JDMessage{}

}

func (this *AlphaController) Post(c mvc.Context, p interface{}) (interface{}, mvc.BingoError) {

	if value,ok:=p.(*JDMessage);ok {
		fmt.Println(value)

		res:=JDMessageResponse{}
		res.Version="1.0"
		res.ShouldEnd=true
		res.Response.Output.Type="PlainText"
		res.Response.Output.Text="欢迎使用！"

		return &res,nil
	}
	return nil,nil
}
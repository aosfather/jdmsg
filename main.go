package main

import "github.com/aosfather/bingo"

func main() {
	app := bingo.TApplication{}
	app.SetHandler(loadService, loadControl)
	app.Run("config.conf")
}

func loadService(context *bingo.ApplicationContext) bool {
	ss := SkillService{}
	ss.aliasPath = context.GetPropertyFromConfig("jdmsg.path")
	ss.msgPath = context.GetPropertyFromConfig("jdmsg.path")
	context.RegisterService("skill", &ss)

	return true
}

func loadControl(mvc *bingo.MvcEngine, context *bingo.ApplicationContext) bool {

	mvc.AddController(&AlphaController{})
	return true
}

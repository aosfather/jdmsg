package main

import "github.com/aosfather/bingo"

func main() {
	app := bingo.TApplication{}
	app.SetHandler(nil, loadControl)
	app.Run("config.conf")
}

func loadControl(mvc *bingo.MvcEngine, context *bingo.ApplicationContext) bool {

	mvc.AddController(&AlphaController{})
	return true
}

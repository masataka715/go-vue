package main

import (
	"gortfolio/config"
	"gortfolio/shiritori"
	"gortfolio/todo"
	"gortfolio/twitter"
	"gortfolio/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	//todo
	todo.Init()
	todo.OnlyGo(router)
	todo.VueAndGo(router)
	//しりとり
	shiritori.OnlyGo(router)
	shiritori.VueAndGo(router)
	// ツイッター
	twitter.Init()
	twitter.Router(router)

	_ = router.Run(":" + config.Config.Port)

}

package main

import (
	"gortfolio/auth"
	"gortfolio/battle"
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
	todo.OnlyGo(router)
	todo.VueAndGo(router)
	//しりとり
	shiritori.OnlyGo(router)
	shiritori.VueAndGo(router)
	// ツイッター
	twitter.Router(router)
	// 認証
	auth.Router(router)
	// 戦闘シーン
	battle.Router(router)

	_ = router.Run(":" + config.Config.Port)

}

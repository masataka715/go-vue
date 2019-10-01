package main

import (
	"gortfolio/config"
	"gortfolio/pkg/auth"
	"gortfolio/pkg/battle"
	"gortfolio/pkg/note"
	"gortfolio/pkg/shiritori"
	"gortfolio/pkg/todo"
	"gortfolio/pkg/twitter"
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
	// ノート
	note.Router(router)

	_ = router.Run(":" + config.Config.Port)

}

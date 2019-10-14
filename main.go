package main

import (
	"go_vue/config"
	"go_vue/pkg/auth"
	"go_vue/pkg/battle"
	"go_vue/pkg/note"
	"go_vue/pkg/shiritori"
	"go_vue/pkg/todo"
	"go_vue/pkg/twitter"
	"go_vue/utils"

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

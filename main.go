package main

import (
	"gortfolio/config"
	"gortfolio/models"
	"gortfolio/shiritori"
	"gortfolio/todo"
	"gortfolio/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	models.Init()
	// Goのみで作ったToDo機能
	todo.OnlyGo(router)
	// VueとGoで作ったToDo機能
	todo.VueAndGo(router)
	//しりとり
	shiritori.Shiritori(router)

	_ = router.Run(":" + config.Config.Port)

}

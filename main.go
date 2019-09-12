package main

import (
	"fmt"
	"gortfolio/config"
	"gortfolio/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	fmt.Println(models.DbConnection)
	data := "test"
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{"data": data})
	})
	// router.POST("/post", func(ctx *gin.Context) {
	// 	ctx.HTML(200, "index.html", gin.H{"data": data})
	// })

	_ = router.Run(":" + config.Config.Port)
}

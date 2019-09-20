package twitter

import (
	"gortfolio/config"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {
	// 一覧表示
	router.GET("/twitter", func(ctx *gin.Context) {
		submitJson(ctx)
	})
	// 新規保存
	router.POST("/twitter", func(ctx *gin.Context) {
		var t Twitter
		err := ctx.BindJSON(&t)
		if err != nil {
			log.Println(err)
		}
		Insert(t.Content)
		submitJson(ctx)
	})
	// 削除
	router.POST("/twitter/delete/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			log.Println(err)
		}
		Delete(id)
		submitJson(ctx)
	})
	// いいね
	router.POST("/twitter/nice/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			log.Println(err)
		}
		AddNice(id)
		submitJson(ctx)
	})
}

func submitJson(ctx *gin.Context) {
	twiiters := GetAll()
	// 別ドメインにアクセスする時に必要
	ctx.Header("Access-Control-Allow-Origin", config.Config.VueUrl)
	ctx.JSON(200, twiiters)
}

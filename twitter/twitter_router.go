package twitter

import (
	"gortfolio/common"
	"gortfolio/common/database"

	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {
	database.Migrate(Twitter{})
	// 一覧表示
	router.GET("/twitter", func(ctx *gin.Context) {
		submitJson(ctx)
	})
	// 新規保存
	router.POST("/twitter", func(ctx *gin.Context) {
		var t Twitter
		common.BindJSON(ctx, &t)
		Insert(t.Content)
		submitJson(ctx)
	})
	// 削除
	router.POST("/twitter/delete/:id", func(ctx *gin.Context) {
		id := common.GetQueryID(ctx)
		Delete(id)
		submitJson(ctx)
	})
	// いいね
	router.POST("/twitter/nice/:id", func(ctx *gin.Context) {
		id := common.GetQueryID(ctx)
		AddNice(id)
		submitJson(ctx)
	})
}

func submitJson(ctx *gin.Context) {
	twiiters := GetAll()
	common.SubmitJson(ctx, twiiters)
}

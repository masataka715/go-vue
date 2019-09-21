package todo

import (
	"gortfolio/common"
	"gortfolio/common/database"

	"github.com/gin-gonic/gin"
)

func VueAndGo(router *gin.Engine) {
	database.Migrate(Todo{})
	// 一覧表示
	router.GET("/todo", func(ctx *gin.Context) {
		submitJson(ctx)
	})
	// 新規保存
	router.POST("/todo", func(ctx *gin.Context) {
		var t Todo
		common.BindJSON(ctx, &t)
		Insert(t.Text, t.Status)
		submitJson(ctx)
	})
	// 編集
	router.POST("/todo/edit/:id", func(ctx *gin.Context) {
		var t Todo
		common.BindJSON(ctx, &t)
		Update(t.ID, t.Text, t.Status)
		submitJson(ctx)
	})
	// 削除
	router.POST("/todo/delete/:id", func(ctx *gin.Context) {
		id := common.GetQueryID(ctx)
		Delete(id)
		submitJson(ctx)
	})
}

func submitJson(ctx *gin.Context) {
	todos := GetAll()
	common.SubmitJson(ctx, todos)
}

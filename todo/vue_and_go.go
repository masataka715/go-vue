package todo

import (
	"gortfolio/config"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func VueAndGo(router *gin.Engine) {
	// 一覧表示
	router.GET("/todo", func(ctx *gin.Context) {
		submitJson(ctx)
	})
	// 新規保存
	router.POST("/todo", func(ctx *gin.Context) {
		var t Todo
		err := ctx.BindJSON(&t)
		if err != nil {
			log.Println(err)
		}
		Insert(t.Text, t.Status)
		submitJson(ctx)
	})
	// 編集
	router.POST("/todo/edit/:id", func(ctx *gin.Context) {
		var t Todo
		err := ctx.BindJSON(&t)
		if err != nil {
			log.Println(err)
		}
		Update(t.ID, t.Text, t.Status)
		submitJson(ctx)
	})
	// 削除
	router.POST("/todo/delete/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		Delete(id)
		submitJson(ctx)
	})
}

func submitJson(ctx *gin.Context) {
	todos := GetAll()
	// 別ドメインにアクセスする時に必要
	ctx.Header("Access-Control-Allow-Origin", config.Config.VueUrl)
	ctx.JSON(200, todos)
}

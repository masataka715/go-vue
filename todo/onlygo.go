package todo

import (
	"gortfolio/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func OnlyGo(router *gin.Engine) {
	//Index
	router.GET("/", func(ctx *gin.Context) {
		todos := models.GetAll()

		lastLetter, _ := ctx.Cookie("lastLetter")
		shiritoriMessage, _ := ctx.Cookie("shiritoriMessage")

		ctx.HTML(200, "index.html", gin.H{
			"todos":            todos,
			"lastLetter":       lastLetter,
			"shiritoriMessage": shiritoriMessage,
		})
	})

	//Create
	router.POST("/new", func(ctx *gin.Context) {
		text := ctx.PostForm("text")
		status := ctx.PostForm("status")
		models.Insert(text, status)
		ctx.Redirect(302, "/")
	})

	//Detail
	router.GET("/detail/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		todo := models.GetOne(id)
		ctx.HTML(200, "detail.html", gin.H{"todo": todo})
	})

	//Update
	router.POST("/update/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		text := ctx.PostForm("text")
		status := ctx.PostForm("status")
		models.Update(id, text, status)
		ctx.Redirect(302, "/")
	})

	//削除確認
	router.GET("/delete_check/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		todo := models.GetOne(id)
		ctx.HTML(200, "delete.html", gin.H{"todo": todo})
	})

	//Delete
	router.POST("/delete/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		models.Delete(id)
		ctx.Redirect(302, "/")
	})
}

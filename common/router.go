package common

import (
	"gortfolio/config"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func BindJSON(ctx *gin.Context, i interface{}) {
	err := ctx.BindJSON(&i)
	if err != nil {
		log.Println(err)
	}
}

func GetQueryID(ctx *gin.Context) int {
	n := ctx.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		log.Println(err)
	}
	return id
}

func SubmitJson(ctx *gin.Context, i interface{}) {
	// 別ドメインにアクセスする時に必要
	ctx.Header("Access-Control-Allow-Origin", config.Config.VueUrl)
	ctx.JSON(200, i)
}
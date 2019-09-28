package common

import (
	"gortfolio/config"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AccessDenyMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.Referer() != config.Config.VueUrl+"/" {
			log.Println(ctx.Request.Referer())
			ctx.JSON(400, gin.H{"err": "Access denied"})
			ctx.Abort()
		}
	}
}

func BindJSON(ctx *gin.Context, i interface{}) {
	err := ctx.BindJSON(&i)
	if err != nil {
		log.Println(err)
	}
}

func GetQueryID(ctx *gin.Context, param_name string) int {
	n := ctx.Param(param_name)
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

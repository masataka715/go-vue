package note

import (
	"go_vue/pkg/common"
	"log"

	"github.com/gin-gonic/gin"
)

func tagRouter(router *gin.Engine) {
	router.POST("/note/tag_create", func(ctx *gin.Context) {
		var tag Tag
		common.BindJSON(ctx, &tag)
		log.Println(tag)
		TagInsert(tag)
		tags := GetAllTags(tag.AuthID)
		common.SubmitJson(ctx, tags)
	})

	router.GET("/note/tag_all/:auth_id", func(ctx *gin.Context) {
		auth_id := common.GetQueryID(ctx, "auth_id")
		tags := GetAllTags(auth_id)
		common.SubmitJson(ctx, tags)
	})
}

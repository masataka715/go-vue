package note

import (
	"gortfolio/common"
	"gortfolio/common/database"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {
	database.Migrate(Note{})

	router.POST("/note/new", func(ctx *gin.Context) {
		var note Note
		common.BindJSON(ctx, &note)
		Insert(note)
		common.SubmitJson(ctx, note)
	})

	router.GET("/note/all/:id", func(ctx *gin.Context) {
		auth_id := common.GetQueryID(ctx)
		// 新規ノート保存待ち
		time.Sleep(10 * time.Millisecond)
		note := GetAllNotes(auth_id)
		common.SubmitJson(ctx, note)
	})

	router.GET("/note/details/:id", func(ctx *gin.Context) {
		note_id := common.GetQueryID(ctx)
		log.Println(note_id)
		note := GetOneNote(note_id)
		common.SubmitJson(ctx, note)
	})
}

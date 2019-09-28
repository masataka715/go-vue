package note

import (
	"gortfolio/common"
	"gortfolio/common/database"
	"time"

	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {
	database.Migrate(Note{})
	router.Use(common.AccessDenyMiddleware())

	router.POST("/note/new", func(ctx *gin.Context) {
		var note Note
		common.BindJSON(ctx, &note)
		NoteInsert(note)
		common.SubmitJson(ctx, note)
	})

	router.GET("/note/all/:auth_id", func(ctx *gin.Context) {
		auth_id := common.GetQueryID(ctx, "auth_id")
		// 新規ノート保存待ち
		time.Sleep(10 * time.Millisecond)
		note := GetAllNotes(auth_id)
		common.SubmitJson(ctx, note)
	})

	router.GET("/note/details/:note_id", func(ctx *gin.Context) {
		note_id := common.GetQueryID(ctx, "note_id")
		note := GetOneNote(note_id)
		common.SubmitJson(ctx, note)
	})

	router.POST("/note/details/", func(ctx *gin.Context) {
		var note Note
		common.BindJSON(ctx, &note)
		NoteUpdate(note)
		common.SubmitJson(ctx, note)
	})

	database.Migrate(Garbage{})

	router.GET("/note/garbage_all/:auth_id", func(ctx *gin.Context) {
		auth_id := common.GetQueryID(ctx, "auth_id")
		garbages := GetGarbageAll(auth_id)
		common.SubmitJson(ctx, garbages)
	})

	router.GET("/note/garbage/:note_id", func(ctx *gin.Context) {
		note_id := common.GetQueryID(ctx, "note_id")
		note := NoteDel(note_id)
		GarbageInsert(note)
		notes := GetAllNotes(note.AuthID)
		common.SubmitJson(ctx, notes)
	})

	router.GET("/note/garbage_details/:garbage_id", func(ctx *gin.Context) {
		garbage_id := common.GetQueryID(ctx, "garbage_id")
		garbage := GetOneGarbage(garbage_id)
		common.SubmitJson(ctx, garbage)
	})

	router.GET("/note/garbage_del/:garbage_id", func(ctx *gin.Context) {
		garbage_id := common.GetQueryID(ctx, "garbage_id")
		GarbageDel(garbage_id)
		// common.SubmitJson(ctx)
	})
}

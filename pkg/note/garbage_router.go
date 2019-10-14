package note

import (
	"go_vue/pkg/common"

	"github.com/gin-gonic/gin"
)

func garbageRouter(router *gin.Engine) {
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

	router.GET("/note/garbage_restore/:garbage_id/:auth_id", func(ctx *gin.Context) {
		garbage_id := common.GetQueryID(ctx, "garbage_id")
		garbage := GetOneGarbage(garbage_id)
		NoteRestore(garbage)
		GarbageDel(garbage_id)

		auth_id := common.GetQueryID(ctx, "auth_id")
		garbages := GetGarbageAll(auth_id)
		common.SubmitJson(ctx, garbages)
	})

	router.GET("/note/garbage_del/:garbage_id/:auth_id", func(ctx *gin.Context) {
		garbage_id := common.GetQueryID(ctx, "garbage_id")
		GarbageDel(garbage_id)
		auth_id := common.GetQueryID(ctx, "auth_id")
		garbages := GetGarbageAll(auth_id)
		common.SubmitJson(ctx, garbages)
	})
}

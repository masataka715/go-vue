package note

import (
	"gortfolio/common/database"
	"time"
)

type Note struct {
	ID        int       `json:"id"`
	AuthID    int       `json:"auth_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func NoteInsert(note Note) {
	db := database.Open()
	db.Create(&Note{AuthID: note.AuthID, Title: note.Title, Content: note.Content})
	db.Close()
}

func NoteUpdate(new_note Note) {
	db := database.Open()
	var note Note
	db.First(&note, new_note.ID)
	note.Title = new_note.Title
	note.Content = new_note.Content
	db.Save(&note)
	db.Close()
}

func NoteDel(id int) Note {
	db := database.Open()
	var note Note
	db.First(&note, id)
	db.Delete(&note)
	db.Close()
	return note
}

func GetAllNotes(auth_id int) []Note {
	db := database.Open()
	var notes []Note
	db.Where("auth_id = ?", auth_id).Order("created_at desc").Find(&notes)
	db.Close()
	return notes
}

func GetOneNote(note_id int) Note {
	db := database.Open()
	var note Note
	db.Where("id = ?", note_id).First(&note)
	db.Close()
	return note
}

type Garbage struct {
	ID        int       `json:"id"`
	AuthID    int       `json:"auth_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func GarbageInsert(note Note) {
	db := database.Open()
	db.Create(&Garbage{AuthID: note.AuthID, Title: note.Title, Content: note.Content})
	db.Close()
}

func NoteRestore(garbage Garbage) {
	db := database.Open()
	db.Create(&Note{AuthID: garbage.AuthID, Title: garbage.Title, Content: garbage.Content})
	db.Close()
}

func GarbageDel(garbage_id int) {
	db := database.Open()
	var garbage Garbage
	db.First(&garbage, garbage_id)
	db.Delete(&garbage)
	db.Close()
}

func GetGarbageAll(auth_id int) []Garbage {
	db := database.Open()
	var garbages []Garbage
	db.Where("auth_id = ?", auth_id).Order("created_at desc").Find(&garbages)
	db.Close()
	return garbages
}

func GetOneGarbage(garbage_id int) Garbage {
	db := database.Open()
	var garbage Garbage
	db.Where("id = ?", garbage_id).First(&garbage)
	db.Close()
	return garbage
}

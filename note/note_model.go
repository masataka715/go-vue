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

func Insert(note Note) {
	db := database.Open()
	db.Create(&Note{AuthID: note.AuthID, Title: note.Title, Content: note.Content})
	db.Close()
}

func Update(new_note Note) {
	db := database.Open()
	var note Note
	db.First(&note, new_note.ID)
	note.Title = new_note.Title
	note.Content = new_note.Content
	db.Save(&note)
	db.Close()
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

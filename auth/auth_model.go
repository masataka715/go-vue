package auth

import "gortfolio/common/database"

type Auth struct {
	ID       int    `json:"id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

//DB追加
func Insert(userName string, password string) {
	db := database.Open()
	db.Create(&Auth{UserName: userName, Password: password})
	db.Close()
}

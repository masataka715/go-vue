package auth

import (
	"gortfolio/common/database"
)

type Auth struct {
	ID           int    `json:"id"`
	EmailAddress string `json:"email_address"`
	Password     string `json:"password"`
}

func Insert(emailAddress string, password string) {
	db := database.Open()
	db.Create(&Auth{EmailAddress: emailAddress, Password: password})
	db.Close()
}

func GetMatchingAuth(auth Auth) Auth {
	db := database.Open()
	newAuth := Auth{}
	db.Where("email_address = ?", auth.EmailAddress).Where("password = ?", auth.Password).First(&newAuth)
	db.Close()
	return newAuth
}

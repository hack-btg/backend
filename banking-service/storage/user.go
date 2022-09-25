package storage

import (
	"encoding/json"
	"log"
	"os"

	"github.com/hack-btg/backend/banking-service/internal/domains/models"
)

type User struct {
	store models.UserInfo
}

func NewUserRepo() *User {
	file, err := os.ReadFile("user_data.json")
	if err != nil {
		log.Println("could not open user data file")
		return nil
	}
	data := models.UserInfo{}
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Println("could not open user data file")
		return nil
	}
	return &User{store: data}
}

func (u *User) Info() models.UserInfo {
	return u.store
}

package storage

import (
	"encoding/json"
	"log"
	"os"

	"github.com/hack-btg/backend/banking-service/internal/domains/models"
)

type Social struct {
	store models.Social
}

func NewSocialRepo() *Social {
	file, err := os.ReadFile("media_data.json")
	if err != nil {
		log.Println("could not open user data file")
		return nil
	}
	data := models.Social{}
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Println("could not open user data file")
		return nil
	}
	return &Social{store: data}
}

func (s *Social) Get() models.Social {
	return s.store
}

package models

type API struct {
	Token    string `json:"token"`
	UserUUID string `json:"user_uuid"`
	UserType string `json:"user_type"`
	Name     string `json:"name"`
}

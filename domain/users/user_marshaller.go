package users

import (
	"encoding/json"
)

type PublicUser struct {
	Id          int64   `json:"id"`
	FirstName   string  `json:"first_name"`
	DateCreated string  `json:"date_created"`
	Status      *string `json:"status"`
}

type PrivateUser struct {
	Id          int64   `json:"id"`
	FirstName   string  `json:"first_name"`
	LastName    string  `json:"last_name"`
	Email       string  `json:"email"`
	DateCreated string  `json:"date_created"`
	Status      *string `json:"status"`
}

type Users []User

func (u *User) Marshall(public bool) interface{} {
	if public {
		return PublicUser{
			Id:          u.Id,
			FirstName:   u.FirstName,
			DateCreated: u.DateCreated,
			Status:      u.Status,
		}
	}

	user, _ := json.Marshal(u) // JSON.stringify
	var privateUser PrivateUser
	json.Unmarshal(user, &privateUser) // JSON.parse
	return privateUser
}

func (us Users) Marshall(public bool) []interface{} {
	users := make([]interface{}, len(us))
	for index, user := range us {
		users[index] = user.Marshall(public)
	}
	return users
}

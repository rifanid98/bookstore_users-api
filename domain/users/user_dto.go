package users

import (
	resp "bookstore_users-api/utils/response"
	"strings"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

func (user *User) Validate() *resp.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return resp.BadRequest("Invalid email address")
	}
	return nil
}

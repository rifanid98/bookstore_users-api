package users

import (
	"bookstore_users-api/datasources/mysql/users_db"
	"bookstore_users-api/utils/dates"
	"bookstore_users-api/utils/errors"
	"fmt"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

	result := usersDB[user.Id]
	if result == nil {
		return errors.NotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

func (user *User) Save() *errors.RestErr {
	currentUser := usersDB[user.Id]
	if currentUser != nil {
		if user.Email == currentUser.Email {
			return errors.BadRequestError(fmt.Sprintf("email %s already registered", user.Email))
		}
		return errors.BadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	}

	user.DateCreated = dates.GetNowString()
	usersDB[user.Id] = user
	return nil
}

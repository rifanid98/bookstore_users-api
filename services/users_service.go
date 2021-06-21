package services

import (
	"bookstore_users-api/domain/users"
	resp "bookstore_users-api/utils/response"
)

func GetUser(userId int64) (*users.User, *resp.RestErr) {
	user := &users.User{
		Id: userId,
	}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return user, nil
}

func CreateUser(user *users.User) (*users.User, *resp.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return user, nil
}

func UpdateUser(user *users.User) (*users.User, *resp.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Update(); err != nil {
		return nil, err
	}

	return user, nil
}

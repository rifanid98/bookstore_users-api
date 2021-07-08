package services

import (
	"bookstore_users-api/domain/users"
	resp "bookstore_users-api/utils/response"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// Implicitly implemented by the usersService struct
type usersServiceInterface interface {
	GetUser(query *users.UserQuery) (users.Users, *resp.RestErr)
	GetUserById(userId int64) (*users.User, *resp.RestErr)
	CreateUser(user *users.User) (*users.User, *resp.RestErr)
	UpdateUser(user *users.User) (*users.User, *resp.RestErr)
	DeleteUser(user *users.User) (*users.User, *resp.RestErr)
	LoginUser(user *users.User) (*users.User, *resp.RestErr)
}

// extends usersServiceInterface and Implements methods according to the contract
// at the UsersServiceInterface, because in go struct it can't be made method in
// the struct itself.
type usersService struct{}

// This variable type struct usersservice and starts in capital that indicates that
// the method is public and can be accessed by other packages
var UsersService usersServiceInterface = &usersService{}

// The implementation of the method of the UsersServiceInterface is made or implemented
// outside the struct separately as below:
func (s *usersService) GetUser(query *users.UserQuery) (users.Users, *resp.RestErr) {
	user := users.User{}
	users, err := user.Find(query.Status)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *usersService) GetUserById(userId int64) (*users.User, *resp.RestErr) {
	user := &users.User{
		Id: userId,
	}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *usersService) CreateUser(user *users.User) (*users.User, *resp.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if user.Password != nil {
		hashed, err := bcrypt.GenerateFromPassword([]byte(*user.Password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println(err.Error())
			return nil, resp.InternalServer("failed to hash the password")
		}

		newPassword := string(hashed)
		user.Password = &newPassword
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *usersService) UpdateUser(user *users.User) (*users.User, *resp.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Update(); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *usersService) DeleteUser(user *users.User) (*users.User, *resp.RestErr) {
	if err := user.Delete(); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *usersService) LoginUser(user *users.User) (*users.User, *resp.RestErr) {
	err := user.GetByCredential()
	if err != nil {
		return nil, err
	}
	return user, nil
}

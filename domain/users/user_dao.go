package users

import (
	"bookstore_users-api/datasources/mysql/users_db"
	"bookstore_users-api/utils/dates"
	e "bookstore_users-api/utils/errors"
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
)

var (
	usersDB = make(map[int64]*User)
)

const (
	emailField = "users_UN"
	emptyRow   = "no rows in result set"
)

func (user *User) Get() *e.RestErr {
	row := users_db.Client.QueryRow(GetUserByIdQuery,
		user.Id,
	)
	err := row.Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.DateCreated,
	)
	if err != nil {
		fmt.Println(err.Error())
		if strings.Contains(err.Error(), emptyRow) {
			return e.NotFoundError(fmt.Sprintf("user with id %d not found", user.Id))
		} else {
			return e.InternalServerError(fmt.Sprintf("failed to get user with id %d", user.Id))
		}
	}
	return nil
}

func (user *User) Save() *e.RestErr {
	user.DateCreated = dates.GetNowString()
	insert, err := users_db.Client.Exec(InsertUserQuery,
		user.FirstName,
		user.LastName,
		user.Email,
		user.DateCreated,
	)

	me, _ := err.(*mysql.MySQLError)
	fmt.Println(me.Message)
	if me != nil && strings.Contains(me.Message, emailField) {
		return e.BadRequestError(fmt.Sprintf("email %s already exists", user.Email))
	}

	id, err := insert.LastInsertId()
	if err != nil {
		return e.InternalServerError("failed to get last insert id")
	}

	user.Id = id
	return nil
}

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

func (user *User) Get() *e.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

	result := usersDB[user.Id]
	if result == nil {
		return e.NotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

func (user *User) Save() *e.RestErr {
	user.DateCreated = dates.GetNowString()
	query := `INSERT INTO users(first_name, last_name, email, date_created)
						VALUES (?,?,?,?);`

	insert, err := users_db.Client.Exec(query,
		user.FirstName,
		user.LastName,
		user.Email,
		user.DateCreated,
	)

	me, _ := err.(*mysql.MySQLError)
	fmt.Println(me.Message)
	if me != nil && strings.Contains(me.Message, "users_UN") {
		return e.BadRequestError(fmt.Sprintf("email %s already exists", user.Email))
	}

	id, err := insert.LastInsertId()
	if err != nil {
		return e.InternalServerError("failed to get last insert id")
	}

	user.Id = id
	return nil
}

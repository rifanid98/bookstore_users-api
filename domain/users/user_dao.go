package users

import (
	"bookstore_users-api/datasources/mysql/users_db"
	"bookstore_users-api/utils/dates"
	e "bookstore_users-api/utils/response"
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func (user *User) Get() *e.RestErr {
	row := users_db.Client.QueryRow(
		GetUserByIdQuery,
		user.Id,
	)
	err := row.Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.DateCreated,
		&user.Email,
		&user.Status,
		&user.Password,
	)
	if err != nil {
		fmt.Println(err.Error())
		if err == sql.ErrNoRows {
			return e.NotFound(fmt.Sprintf("user with id %d not found", user.Id))
		} else {
			return e.InternalServer(fmt.Sprintf("failed to get user with id %d", user.Id))
		}
	}
	return nil
}

func (user *User) Find(status string) (Users, *e.RestErr) {
	var rows *sql.Rows
	var err error

	if len(status) > 0 {
		rows, err = users_db.Client.Query(
			GetUserByStatusQuery,
			status,
		)
	} else {
		rows, err = users_db.Client.Query(GetUserQuery)
	}

	if err != nil {
		fmt.Println(err.Error())
		return nil, e.InternalServer("failed to get users")
	}

	users := make(Users, 0)
	for rows.Next() {
		err := rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.DateCreated,
			&user.Email,
			&user.Status,
			&user.Password,
		)
		if err != nil {
			fmt.Println(err.Error())
			return nil, e.InternalServer("failed to fetch map from database")
		}
		users = append(users, *user)
	}

	return users, nil
}

func (user *User) Save() *e.RestErr {
	user.DateCreated = dates.GetNowString()
	result, err := users_db.Client.Exec(
		InsertUserQuery,
		user.FirstName,
		user.LastName,
		user.Email,
		user.DateCreated,
		user.Status,
		user.Password,
	)

	me, _ := err.(*mysql.MySQLError)
	if me != nil && me.Number == 1062 {
		return e.BadRequest(fmt.Sprintf("email %s already exists", user.Email))
	}

	id, err := result.LastInsertId()
	if err != nil {
		return e.InternalServer("failed to get last insert id")
	}

	user.Id = id
	return nil
}

func (user *User) Update() *e.RestErr {
	_, err := users_db.Client.Exec(
		UpdateUserQuery,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Id,
	)

	if err != nil {
		fmt.Println(err.Error())
		return e.InternalServer(fmt.Sprintf("failed to update user with id %d", user.Id))
	}

	return nil
}

func (user *User) Delete() *e.RestErr {
	result, err := users_db.Client.Exec(
		DeleteUserQuery,
		user.Id,
	)

	if err != nil {
		return e.InternalServer(fmt.Sprintf("failed to delete user with id %d", user.Id))
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return e.InternalServer(fmt.Sprintf("failed to delete user with id %d", user.Id))
	}
	if affectedRows < 1 {
		return e.NotFound(fmt.Sprintf("user with id %d not found", user.Id))
	}

	return nil
}

package users

import (
	"bookstore_users-api/datasources/mysql/users_db"
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/rifanid98/bookstore_utils-go/dates"
	"github.com/rifanid98/bookstore_utils-go/logger"
	resp "github.com/rifanid98/bookstore_utils-go/response"
	"golang.org/x/crypto/bcrypt"
)

func (user *User) Get() *resp.RestErr {
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
		logger.Error("Error when scanning data from the database", err)
		if err == sql.ErrNoRows {
			return resp.NotFound(fmt.Sprintf("user with id %d not found", user.Id))
		} else {
			return resp.InternalServer(fmt.Sprintf("Database Error. failed to get user with id %d", user.Id))
		}
	}
	return nil
}

func (user *User) Find(status string) (Users, *resp.RestErr) {
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
		logger.Error("error when retrieving data from the database", err)
		fmt.Println(err.Error())
		return nil, resp.InternalServer("Database Error. failed to get users")
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
			logger.Error("Error when scanning data from the database", err)
			fmt.Println(err.Error())
			return nil, resp.InternalServer("Logic Error. failed to fetch map from database")
		}
		users = append(users, *user)
	}

	return users, nil
}

func (user *User) Save() *resp.RestErr {
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
		return resp.BadRequest(fmt.Sprintf("email %s already exists", user.Email))
	}

	id, err := result.LastInsertId()
	if err != nil {
		return resp.InternalServer("failed to get last insert id")
	}

	user.Id = id
	return nil
}

func (user *User) Update() *resp.RestErr {
	_, err := users_db.Client.Exec(
		UpdateUserQuery,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Id,
	)

	if err != nil {
		fmt.Println(err.Error())
		return resp.InternalServer(fmt.Sprintf("failed to update user with id %d", user.Id))
	}

	return nil
}

func (user *User) Delete() *resp.RestErr {
	result, err := users_db.Client.Exec(
		DeleteUserQuery,
		user.Id,
	)

	if err != nil {
		return resp.InternalServer(fmt.Sprintf("failed to delete user with id %d", user.Id))
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return resp.InternalServer(fmt.Sprintf("failed to delete user with id %d", user.Id))
	}
	if affectedRows < 1 {
		return resp.NotFound(fmt.Sprintf("user with id %d not found", user.Id))
	}

	return nil
}

func (user *User) GetByCredential() *resp.RestErr {
	password := []byte(*user.Password)
	row := users_db.Client.QueryRow(
		GetUserByCredential,
		user.Email,
	)

	if err := row.Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.DateCreated,
		&user.Email,
		&user.Status,
		&user.Password,
	); err != nil {
		if err == sql.ErrNoRows {
			return resp.NotFound("user not found")
		}
		logger.Error("Error when scanning data from the database", err)
		fmt.Println(err.Error())
		return resp.InternalServer("Logic Error. failed to fetch map from database")
	}

	err := bcrypt.CompareHashAndPassword([]byte(*user.Password), password)
	if err == bcrypt.ErrMismatchedHashAndPassword {
		logger.Info("User gives wrong credentials")
		return resp.Unauthorized("Invalid credentials")
	}

	return nil
}

package users

const (
	GetUserByIdQuery = "SELECT * FROM users WHERE id=?"
	InsertUserQuery  = "INSERT INTO users(first_name, last_name, email, date_created) VALUES (?,?,?,?);"
	UpdateUserQuery  = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
)

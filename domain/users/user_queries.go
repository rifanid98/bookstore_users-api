package users

const (
	GetUserQuery         = "SELECT * FROMa users"
	GetUserByIdQuery     = "SELECT * FROM users WHERE id = ?"
	InsertUserQuery      = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES (?,?,?,?,?,?);"
	UpdateUserQuery      = "UPDATE users SET first_name = ?, last_name = ?, email = ? WHERE id = ?;"
	DeleteUserQuery      = "DELETE FROM users WHERE id = ?"
	GetUserByStatusQuery = "SELECT * FROM users WHERE status = ?"
	GetUserByCredential  = "SELECT * FROM users WHERE email=?"
)

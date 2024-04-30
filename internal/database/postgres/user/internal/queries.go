package queries

const (
	CREATE     = "INSERT INTO users (email, password) VALUES ($1, $2) RETURNING *;"
	GetByEmail = "SELECT * FROM users WHERE email = $1;"
)

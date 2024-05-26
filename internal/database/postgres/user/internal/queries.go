package queries

const (
	CREATE     = "INSERT INTO users (email, password, role) VALUES ($1, $2, $3) RETURNING *;"
	GetByEmail = "SELECT * FROM users WHERE email = $1;"
)

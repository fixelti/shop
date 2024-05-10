package queries

const (
	CREATE = "INSERT INTO products(name, description, price, image_url) VALUES ($1, $2, $3, $4) RETURNING *;"
)

package queries

const (
	CREATE  = "INSERT INTO products(name, description, price, image_url) VALUES ($1, $2, $3, $4) RETURNING *;"
	GetByID = "SELECT * FROM products WHERE id = $1;"
	GetList = "SELECT * FROM products LIMIT $1 OFFSET $2;"
)

package userrepository

const (
	GetUsersQuery = "SELECT id, email, first_name, last_name, created_at, updated_at FROM users"
	AddUsersQuery = "INSERT INTO users (email, first_name, last_name) VALUES ($1, $2, $3)"
)

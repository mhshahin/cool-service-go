package userrepository

const (
	GetUsersQuery = "SELECT id, email, name, created_at, updated_at FROM users"
	AddUsersQuery = "INSERT INTO users (email, name) VALUES ($1, $2)"
)

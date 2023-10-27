package model

import (
	"fmt"
	"time"
)

type User struct {
	ID        int       `json:"id,omitempty"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (u User) GetFullname() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

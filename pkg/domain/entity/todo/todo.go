package todo

import "github.com/alifudin-a/golang-todo-app/pkg/domain/helper"

// User : entity model for user
type User struct {
	ID int `json:"id,omitempty" db:"id"`
	Username string `json:"username,omitempty" db:"username"`
	Password string `json:"password,omitempty" db:"password"`
	FullName string `json:"fullname,omitempty" db:"fullname"`
	Email string `json:"email,omitempty" db:"email"`
}

// Todo : entity model for todo
type Todo struct {
	ID int `json:"id,omitempty" db:"id"`
	Title string `json:"title,omitempty" db:"title"`
	Description string `json:"description,omitempty" db:"description"`
	OwnerID int `json:"owner_id,omitempty" db:"owner_id"`
	CreatedAt helper.NullString `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt helper.NullString `json:"updated_at,omitempty" db:"updated_at"`
}
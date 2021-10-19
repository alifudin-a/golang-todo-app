package todo

type User struct {
	ID int `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	FullName string `json:"fullname" db:"fullname"`
	Email string `json:"email" db:"email"`
}

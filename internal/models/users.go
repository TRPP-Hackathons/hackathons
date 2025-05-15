package models

type User struct {
	ID        int    `json:"id" db:"id"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	ImageURL  string `json:"image_url" db:"image_url"`
}

type Participant struct {
	User
	Role string `json:"role" db:"role"`
}

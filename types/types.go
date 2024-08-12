package types

import "time"


type User struct {
	ID                int       `json:"id"`
	FirstName         string    `json:"firstName"`
	LastName          string    `json:"lastName"`
	Username          string    `json:"username,omitempty"`
	Email             string    `json:"email"`
	Password          string    `json:"-"`
	PasswordChangedAt time.Time `json:"passwordChangedAt,omitempty"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
	DeletedAt         time.Time `json:"deletedAt,omitempty"`
}

type RegisterPayload struct {
	FirstName string
	LastName  string
	Email     string
	Username  string
	Password  string
	// passwordChangedAt date
	// deletedAt date
}

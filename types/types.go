package types

import "time"

type Response struct {
	Code int `json:"code,omitempty"`
	// Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

type UserRepo interface {
	GetUserByEmail(email string) (*User, error)
	CreateUser(user User) error
}

type User struct {
	ID                string     `json:"id"`
	FirstName         string     `json:"firstName"`
	LastName          string     `json:"lastName"`
	Username          *string    `json:"username,omitempty"`
	Email             string     `json:"email"`
	Password          string     `json:"password"`
	PasswordChangedAt *time.Time `json:"passwordChangedAt,omitempty"`
	CreatedAt         time.Time  `json:"createdAt"`
	UpdatedAt         time.Time  `json:"updatedAt"`
	DeletedAt         *time.Time `json:"deletedAt,omitempty"`
}

type RegisterPayload struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Username  string `json:"username,omitempty"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8"`
}

type LoginPayload struct {
	Email    string `json:"email" valid:"required~please provide an email,email"`
	Password string `json:"password" valid:"required~please provide a password"`
}
type LoginResponse struct {
	Token string `json:"token"`
	User  *User  `json:"user"`
}

package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bsin:"id"`
	FirstName    *string            `json:"firstname" validate:"required, min=2, max=100"`
	LastName     *string            `json:"lastname"`
	Password     *string            `json:"password" validate:"required, min=8"`
	Email        *string            `json:"email"`
	Phone        *string            `json:"phone"`
	Token        *string            `json:"token"`
	UserType     *string            `json:"usertype"`
	RefreshToken *string            `json:"refreshtoken"`
	CreatedAt    time.Time          `json:"createdat"`
	UpdatedAt    time.Time          `json:"updatedat"`
	UserID       string
}

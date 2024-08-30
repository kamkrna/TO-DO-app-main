package pkg

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserDB struct{
	ID primitive.ObjectID `bson:"_id,omitempty"`
	Username string `json:"username" validate:"required,min=2,max=50"`
	Password string `json:"password" validate:"required,min=8"`
	// Email string `json:"email" validate:"email,required"`
}

type TaskDB struct{
	ID primitive.ObjectID `bson:"_id,omitempty"`
	UserId primitive.ObjectID `bson:"user_id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Due_date time.Time `bson:"due_date"`
	// Status string `json:"status"`
}
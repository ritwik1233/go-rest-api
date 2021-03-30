package models

import "time"

type SessionCollection struct {
	ID          string    `bson:"_id"`
	Value       string    `bson:"value"`
	Email       string    `bson:"email"`
	CreatedDate time.Time `bson:"createdDate"`
}

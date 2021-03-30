package models

import "time"

type MessageCollection struct {
	ID          string    `bson:"_id"`
	Message     string    `bson:"message"`
	CreatedBy   string    `bson:"createdBy"`
	CreatedDate time.Time `bson:"createdDate"`
}

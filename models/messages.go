package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageCollection struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Message     string             `json:"message,omitempty" bson:"message,omitempty"`
	CreatedBy   string             `json:"createdBy,omitempty" bson:"createdBy,omitempty"`
	CreatedDate time.Time          `json:"createdDate,omitempty" bson:"createdDate,omitempty"`
}

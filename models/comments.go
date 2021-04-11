package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentsCollection struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Comment     string             `json:"comment,omitempty" bson:"comment"`
	MessageId   primitive.ObjectID `json:"messageId,omitempty" bson:"messageId"`
	CreatedBy   string             `json:"createdBy,omitempty" bson:"createdBy"`
	CreatedDate time.Time          `json:"createdDate,omitempty" bson:"createdDate"`
}

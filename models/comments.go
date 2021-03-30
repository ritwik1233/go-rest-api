package models

import "time"

type CommentsCollection struct {
	ID          string    `bson:"_id"`
	Comment     string    `bson:"comment"`
	MessageId   string    `bson:"messageId"`
	CreatedBy   string    `bson:"createdBy"`
	CreatedDate time.Time `bson:"createdDate"`
}

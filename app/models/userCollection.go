package models

type UserCollection struct {
	ID       string `bson:"_id"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
	Username string `bson:"username"`
}

package dev

type Keys struct {
	MongoURI   string
	PORT       string
	SessionKey string
}

func (k *Keys) Initialize() {
	(*k).MongoURI = "mongodb://localhost:9001"
	(*k).PORT = "3000"
}

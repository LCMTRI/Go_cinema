package model

type Movie struct {
	Id       string   `json:"_id" bson:"_id"`
	Isbn     string   `json:"name" bson:"name"`
	Title    string   `json:"email" bson:"email"`
	Director Director `json:"director" bson:"director"`
}

type Director struct {
	Firstname string `json:"firstname" bson:"firstname"`
	Lastname  string `json:"lastname" bson:"lastname"`
}

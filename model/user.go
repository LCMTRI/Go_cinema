package model

type User struct {
	Id             string   `json:"_id" bson:"_id"`
	Name           string   `json:"name" bson:"name"`
	Email          string   `json:"email" bson:"email"`
	Password       string   `json:"password" bson:"password"`
	Watched_movies []string `json:"watched_movies" bson:"watched_movies"`
}

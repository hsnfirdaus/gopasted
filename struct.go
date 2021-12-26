package main
import (
	"time"
)

type Error struct {
	Error bool `json:"error"`
	Message string `json:"message"`
}

type CurrentPage struct {
	Page int64 `json:"page"`
}

type Pasted struct {
	ID string `json:"id" bson:"_id"`
	Title string `json:"title" bson:"title"`
	Syntax string `json:"syntax" bson:"syntax"`
	Public bool `json:"public" bson:"public"`
	Created time.Time `json:"created" bson:"created"`
	Protected bool `json:"protected"`
}

type PastedWithPassword struct {
	ID string `bson:"_id"`
	Title string `bson:"title"`
	Syntax string `bson:"syntax"`
	Public bool `bson:"public"`
	Created time.Time `json:"created" bson:"created"`
	Password string `bson:"password"`
}

type PasswordParam struct{
	Password string `json:"password" bson:"password"`
}
type PastedContent struct{
	Content string `json:"content"`
}
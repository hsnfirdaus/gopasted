package main

import (
	"net/http"
	"context"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func DetailHandler(w http.ResponseWriter, r *http.Request){
	var id string
	var err error

	id = mux.Vars(r)["id"]

	if id == "" {
		Response(w,Error{
			Error: true,
			Message: "ID Tidak Valid!",
		})
		return
	}

	collection := GetDB().Collection("pasted")

	cursor := collection.FindOne(context.TODO(),bson.M{"_id":id})
	if cursor == nil {
		Response(w,Error{
			Error: true,
			Message: "Gagal mendapatkan hasil!",
		})
		return
	}

	var pasted PastedWithPassword
	err = cursor.Decode(&pasted)
	if err != nil {
		Response(w,Error{
			Error: true,
			Message: "Gagal decode data!",
		})
		return
	}
	var protected bool
	if pasted.Password != "" {
		protected =true
	} else {
		protected =false
	}
	Response(w,Pasted{
		ID:pasted.ID,
		Title:pasted.Title,
		Syntax:pasted.Syntax,
		Public:pasted.Public,
		Created:pasted.Created,
		Protected:protected,
	})
	return
}
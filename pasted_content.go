package main

import (
	"net/http"
	"context"
	"encoding/json"
	"os"
	"fmt"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func ContentHandler(w http.ResponseWriter, r *http.Request){
	var id string
	var param PasswordParam
	var err error

	id = mux.Vars(r)["id"]

	if id == "" {
		Response(w,Error{
			Error: true,
			Message: "ID Tidak Valid!",
		})
		return
	}

	err = json.NewDecoder(r.Body).Decode(&param)

	if err != nil {
		param.Password = ""
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
	
	if pasted.Password == param.Password {
		data, err := os.ReadFile(fmt.Sprintf("content/%s.dat",id))
		if err != nil {
			Response(w,Error{
				Error: true,
				Message: "Konten hilang!",
			})
			return
		}
		Response(w,PastedContent{
			Content:string(data),
		})
		return
	}

	Response(w,Error{
		Error: true,
		Message: "Password Salah!",
	})

	return
}
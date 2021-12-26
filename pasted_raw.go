package main

import (
	"net/http"
	"context"
	"os"
	"fmt"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func RawHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Content-Type","text/plain")

	var id string
	var param PasswordParam
	var err error

	id = mux.Vars(r)["id"]

	if id == "" {
		w.Write([]byte("ID TIDAK VALID!"))
		return
	}

	param.Password = r.FormValue("password")

	if err != nil {
		param.Password = ""
	}

	collection := GetDB().Collection("pasted")

	cursor := collection.FindOne(context.TODO(),bson.M{"_id":id})
	if cursor == nil {
		w.Write([]byte("Gagal Mendapatkan Hasil!"))
		return
	}

	var pasted PastedWithPassword
	err = cursor.Decode(&pasted)
	if err != nil {
		w.Write([]byte("Gagal decode data!"))
		return
	}
	
	if pasted.Password == param.Password {
		data, err := os.ReadFile(fmt.Sprintf("content/%s.dat",id))
		if err != nil {
			w.Write([]byte("Konten Hilang!"))
			return
		}
		w.Write([]byte(data))
		return
	}

	w.Write([]byte("Password Salah!"))

	return
}
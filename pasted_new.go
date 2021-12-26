package main

import (
	"time"
	"log"
	"fmt"
	"os"
	"context"
	"net/http"
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
)

type NewPasted struct {
	Title string `json:"title"`
	Content string `json:"content"`
	Syntax string `json:"syntax"`
	Public bool `json:"public"`
	Password string `json:"password"`
}

type SuccessPasted struct {
	ID interface{} `json:"id"`
}

func NewHandler(w http.ResponseWriter, r *http.Request){

	var pasted NewPasted
	var response interface{}

	err := json.NewDecoder(r.Body).Decode(&pasted)

	if err != nil {
		response = Error{
			Error:true,
			Message:"Request tidak valid, pastikan tidak ada yang kosong!",
		}
		Response(w,response)
		return
	}

	if len(pasted.Content) < 10 {
		response = Error{
			Error:true,
			Message:"Konten minimal 10 karakter!",
		}
		Response(w,response)
		return
	}

	if len(pasted.Content) > 5242880 {
		response = Error{
			Error:true,
			Message:"Konten Maksimal 5242880 karakter!",
		}
		Response(w,response)
		return
	}

	randomId := RandomID()

	collection := GetDB().Collection("pasted")

	_, err = collection.InsertOne(context.TODO(), bson.M{
		"_id":randomId,
		"title":pasted.Title,
		"syntax":pasted.Syntax,
		"public":pasted.Public,
		"password":pasted.Password,
		"created":time.Now(),
	})
	if err != nil {
		response = Error{
			Error:true,
			Message:"Gagal menambahkan data ke database!",
		}
		Response(w,response)
		return
	}


	err = os.WriteFile(fmt.Sprintf("content/%s.dat",randomId),[]byte(pasted.Content),0644)

	if err != nil {
		log.Fatal(err)
		_, err = collection.DeleteOne(context.TODO(), bson.M{
			"_id":randomId,
		})
		if err != nil {
			response = Error{
				Error:true,
				Message:"Gagal menyimpan pasted file dan menghapus db!",
			}
			Response(w,response)
			return

		}
		response = Error{
			Error:true,
			Message:"Gagal menyimpan pasted file!",
		}
		Response(w,response)
		return
	}

	response = SuccessPasted{
		ID:randomId,
	}
	Response(w,response)
	return

}
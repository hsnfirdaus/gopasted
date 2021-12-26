package main

import (
	"net/http"
	"context"
	"strconv"
	"math"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ResponseList struct {
	Total float64 `json:"total"`
	Result interface{} `json:"result"`
}

func ListHandler(w http.ResponseWriter, r *http.Request){
	var pastedList []Pasted
	var err error
	var limit int64 = 10
	var page int64

	page, err = strconv.ParseInt(r.FormValue("page"), 10, 64)

	if err != nil || page == 0 {
		page = 1
	}

	start := ( page - 1 ) * limit

	collection := GetDB().Collection("pasted")

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"created",-1}})
	findOptions.SetLimit(limit)
	findOptions.SetSkip(start)


	total, err := collection.CountDocuments(context.TODO(),bson.M{"public":true})
	cursor, err := collection.Find(context.TODO(),bson.M{"public":true}, findOptions)
	if err != nil {
		Response(w,Error{
			Error: true,
			Message: "Gagal mendapatkan hasil!",
		})
		return
	}

	for cursor.Next(context.TODO()) {
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
		pastedList = append(pastedList, Pasted{
			ID:pasted.ID,
			Title:pasted.Title,
			Syntax:pasted.Syntax,
			Public:pasted.Public,
			Created:pasted.Created,
			Protected:protected,
		})
	}
	Response(w,ResponseList{
		Total:math.Ceil( float64(total) / float64(limit) ),
		Result:pastedList,
	})
	return
}
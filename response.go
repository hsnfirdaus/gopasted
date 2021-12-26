package main

import (
	"encoding/json"
	"net/http"
)
func Response(w http.ResponseWriter, response interface{}){
	json.NewEncoder(w).Encode(response)
}
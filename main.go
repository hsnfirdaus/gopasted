package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main(){
	log.Println("MEMULAI APLIKASI PASTED")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Tidak dapat meload file .env")
	}
	
	r := mux.NewRouter()

	apiRouter := r.PathPrefix("/api").Subrouter()
	apiRouter.Use(globalMiddleware)

	apiRouter.HandleFunc("/pasted",NewHandler).Methods("POST")

	apiRouter.HandleFunc("/pasted",ListHandler).Methods("GET")

	apiRouter.HandleFunc("/pasted/{id}",DetailHandler).Methods("GET")

	apiRouter.HandleFunc("/pasted/{id}/content",ContentHandler)

	r.HandleFunc("/raw/{id}",RawHandler)

	log.Fatal(http.ListenAndServe(":80",r))
}
func globalMiddleware(next http.Handler) (http.Handler) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		w.Header().Add("Content-Type","application/json")
		next.ServeHTTP(w,r)
	})
}
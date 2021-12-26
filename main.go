package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)
type spaHandler struct {
	staticPath string
	indexPath  string
}

// ServeHTTP inspects the URL path to locate a file within the static dir
// on the SPA handler. If a file is found, it will be served. If not, the
// file located at the index path on the SPA handler will be served. This
// is suitable behavior for serving an SPA (single page application).
func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // get the absolute path to prevent directory traversal
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
        // if we failed to get the absolute path respond with a 400 bad request
        // and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

    // prepend the path with the path to the static directory
	path = filepath.Join(h.staticPath, path)

    // check whether a file exists at the given path
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// file does not exist, serve index.html
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
        // if we got an error (that wasn't that the file doesn't exist) stating the
        // file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    // otherwise, use http.FileServer to serve the static dir
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

func main(){
	log.Println("MEMULAI APLIKASI PASTED")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Tidak dapat meload file .env")
	}
	
	r := mux.NewRouter()

	mux.CORSMethodMiddleware(r)

	apiRouter := r.PathPrefix("/api").Subrouter()
	apiRouter.Use(globalMiddleware)

	apiRouter.HandleFunc("/{any:.+}",optionsHandler).Methods("OPTIONS");

	apiRouter.HandleFunc("/pasted",NewHandler).Methods("POST")

	apiRouter.HandleFunc("/pasted",ListHandler).Methods("GET")

	apiRouter.HandleFunc("/pasted/{id}",DetailHandler).Methods("GET")

	apiRouter.HandleFunc("/pasted/{id}/content",ContentHandler)

	r.HandleFunc("/raw/{id}",RawHandler)

	spa := spaHandler{staticPath: "build", indexPath: "index.html"}
	r.PathPrefix("/").Handler(spa)

	log.Fatal(http.ListenAndServe(":80",r))
}
func globalMiddleware(next http.Handler) (http.Handler) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		w.Header().Add("Content-Type","application/json")
		w.Header().Add("Access-Control-Allow-Headers","content-type, accept")
		w.Header().Add("Access-Control-Allow-Origin","*")
		next.ServeHTTP(w,r)
	})
}

func optionsHandler(_ http.ResponseWriter, _ *http.Request){
	return
}
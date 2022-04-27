package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandlerRequest() {
	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"content-type"}),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowCredentials(),
	)
	var dir string
	//http.FileServer(http.Dir("./web/assets/"))
	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))
	r.Use(cors)
	r.HandleFunc("/addFile", AddFile).Methods("POST", "OPTIONS")
	r.HandleFunc("/download", GetImage).Methods(http.MethodGet, "OPTIONS")

	log.Fatal(http.ListenAndServe(":10002", r))
}

func main() {
	HandlerRequest()
	//files, _ := filepath.Glob("*.pdf")
	//fmt.Printf("%q\n", files)
}

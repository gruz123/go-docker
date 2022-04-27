package main

import (
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

	//r.HandleFunc("/list", Project_list).Methods("POST", "OPTIONS")
	//r.HandleFunc("/addProject", AddProject).Methods("POST", "OPTIONS")
	//r.HandleFunc("/addlistcode", AddListCode).Methods("POST", "OPTIONS")
	//r.HandleFunc("/listcode", ListCode).Methods("POST", "OPTIONS")
	//r.HandleFunc("/getdatacode", GetData).Methods("POST", "OPTIONS")
	//r.HandleFunc("/pushparamcode", Pushdate).Methods("POST", "OPTIONS")

	r.HandleFunc("/addFile", AddFile).Methods("POST", "OPTIONS")
	r.HandleFunc("/download", GetImage).Methods(http.MethodGet, "OPTIONS")
	//r.HandleFunc("/deleteCode", DeleteCode).Methods("POST", "OPTIONS")
	////r.HandleFunc("/chaingeName", ChaingeName).Methods("POST", "OPTIONS")
	//r.HandleFunc("/deleteApp", DeleteApp).Methods("POST", "OPTIONS")
	//r.HandleFunc("/getName", GetName).Methods("POST", "OPTIONS")

	log.Fatal(http.ListenAndServe(":10002", r))
}

func main() {
	HandlerRequest()
	//files, _ := filepath.Glob("*.pdf")
	//fmt.Printf("%q\n", files)
}

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"net/http"
)

type ListName struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func Project_list(w http.ResponseWriter, r *http.Request) {
	connStr := "user=postgres password=0208 dbname=AppData sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from nameapp")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	listname := []ListName{}

	for rows.Next() {
		p := ListName{}
		err := rows.Scan(&p.Id, &p.Name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		listname = append(listname, p)
	}
	err = json.NewEncoder(w).Encode(listname)
}
func AddProject(w http.ResponseWriter, r *http.Request) {
	var decoded interface{}

	//читаем тело запроса
	body, err1 := ioutil.ReadAll(r.Body)
	if err1 != nil {
		log.Println("R.body err")
	}
	err := json.Unmarshal(body, &decoded)
	if err != nil {
		log.Println(err)
	}

	connStr := "user=postgres password=0208 dbname=AppData sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	result, err := db.Exec("insert into nameapp (name) values ($1)",
		decoded)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.LastInsertId()) // не поддерживается
	fmt.Println(result.RowsAffected()) // количество добавленных строк
}

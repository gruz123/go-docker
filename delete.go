package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Name struct {
	Name string `Name:"id"`
}

func DeleteCode(w http.ResponseWriter, r *http.Request) {
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
	// удаляем строку с id=2
	result, err := db.Exec("delete from listappcode where id = $1", decoded)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.RowsAffected()) // количество удаленных строк
}
func ChaingeName(w http.ResponseWriter, r *http.Request) {
	var decoded []interface{}
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
	result, err := db.Exec("update nameapp set name = $1 where id = $2", decoded[0], decoded[1])
	if err != nil {
		panic(err)
	}
	fmt.Println(result.RowsAffected()) // количество обновленных строк
}
func DeleteApp(w http.ResponseWriter, r *http.Request) {
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
	// удаляем строку с id=2
	result, err := db.Exec("delete from nameapp where id = $1", decoded)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.RowsAffected()) // количество удаленных строк
}
func GetName(w http.ResponseWriter, r *http.Request) {

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
	rows, err := db.Query("select name from nameapp where id=$1", decoded)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	listcode := []Name{}

	for rows.Next() {
		p := Name{}
		err := rows.Scan(&p.Name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		listcode = append(listcode, p)
	}
	err = json.NewEncoder(w).Encode(listcode)
}

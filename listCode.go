package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type CodeValu struct {
	Id     int    `json:"id"`
	IdJson string `json:"idJson"`
}

//func AddListCode(w http.ResponseWriter, r *http.Request) {
//	var decoded interface{}
//	//читаем тело запроса
//	body, err1 := ioutil.ReadAll(r.Body)
//	if err1 != nil {
//		log.Println("R.body err")
//	}
//	err := json.Unmarshal(body, &decoded)
//	if err != nil {
//		log.Println(err)
//	}
//	//rand.Seed(time.Now().UnixNano())
//	//digits := "123456789"
//	////specials := "~=+%^*/()[]{}/!@#$?|"
//	//all := "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
//	//	"abcdefghijklmnopqrstuvwxyz" +
//	//	digits
//	//length := 16
//	//buf := make([]byte, length)
//	//buf[0] = digits[rand.Intn(len(digits))]
//	//for i := 2; i < length; i++ {
//	//	buf[i] = all[rand.Intn(len(all))]
//	//}
//	//rand.Shuffle(len(buf), func(i, j int) {
//	//	buf[i], buf[j] = buf[j], buf[i]
//	//})
//	//str := string(buf) // Например "3i[g0|)z"
//	//idvalu := strings.Replace(str, "\u0000", "", -1)
//
//	connStr := "user=postgres password=0208 dbname=AppData sslmode=disable"
//	db, err := sql.Open("postgres", connStr)
//	if err != nil {
//		panic(err)
//	}
//	defer db.Close()
//	//result, err := db.Exec("insert into listappcode (idJson, idapp) values ($1,$2)", idvalu, decoded)
//	//if err != nil {
//	//	panic(err)
//	//}
//	//fmt.Println(result.LastInsertId()) // не поддерживается
//	//fmt.Println(result.RowsAffected()) // количество добавленных строк
//}
func ListCode(w http.ResponseWriter, r *http.Request) {
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
	rows, err := db.Query("select id,idjson from listappcode where idapp=$1", decoded)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	listcode := []CodeValu{}

	for rows.Next() {
		p := CodeValu{}
		err := rows.Scan(&p.Id, &p.IdJson)
		if err != nil {
			fmt.Println(err)
			continue
		}
		listcode = append(listcode, p)
	}
	err = json.NewEncoder(w).Encode(listcode)
}

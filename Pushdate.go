package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Pushdate(w http.ResponseWriter, r *http.Request) {
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
	log.Println(decoded[10])
	connStr := "user=postgres password=0208 dbname=AppData sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	result, err := db.Exec("UPDATE listappcode SET media_sourse=$1,media_sourse_clicadila=$2,country=$3,creative_text=$4,site_name=$5,site_url=$6,appstore_page=$7,onboard=$8,add_description=$9, jsonapp=$11 WHERE id =$10", decoded[0], decoded[1], decoded[2], decoded[3], decoded[4], decoded[5], decoded[6], decoded[7], decoded[8], decoded[9], decoded[10])
	if err != nil {
		panic(err)
	}
	fmt.Println(result.LastInsertId()) // не поддерживается
	fmt.Println(result.RowsAffected()) // количество добавленных строк
}

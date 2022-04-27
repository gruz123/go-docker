package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type listcodelist struct {
	MediaSourse           interface{} `json:"media_sourse"`
	MediaSourseClicadlila interface{} `json:"media_sourse_clicadila"`
	Country               interface{} `json:"country"`
	Creative_text         interface{} `json:"creative_text"`
	Site_name             interface{} `json:"site_name"`
	Site_url              interface{} `json:"site_url"`
	Appstore_page         interface{} `json:"appstore_page"`
	Onboard               interface{} `json:"onboard"`
	Add_description       interface{} `json:"add_on"`
	JsonApp               interface{} `json:"Jsondescription"`
	Idjson                interface{} `json:"idjsApp"`
}

func GetData(w http.ResponseWriter, r *http.Request) {

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
	rows, err := db.Query("select media_sourse,media_sourse_clicadila,country,creative_text,site_name,site_url,appstore_page,onboard,add_description,idjson,jsonapp from listappcode where id=$1", decoded)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	listcode := []listcodelist{}

	for rows.Next() {
		p := listcodelist{}
		err := rows.Scan(&p.MediaSourse, &p.MediaSourseClicadlila, &p.Country, &p.Creative_text, &p.Site_name, &p.Site_url, &p.Appstore_page, &p.Onboard, &p.Add_description, &p.Idjson, &p.JsonApp)
		if err != nil {
			fmt.Println(err)
			continue
		}
		listcode = append(listcode, p)
	}
	err = json.NewEncoder(w).Encode(listcode)
}

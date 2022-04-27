package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func getImageFromFilePath(filePath string) ([]byte, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	readAll, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return readAll, err
}

func GetImage(w http.ResponseWriter, r *http.Request) {
	var imgPath string
	imgPath = r.Header.Get("img")
	imgPath = fmt.Sprintf("./%s", imgPath)
	path, err := getImageFromFilePath(imgPath)
	if err != nil {
		log.Println("shit happens")
	}
	w.Write(path)
	log.Println("Хуита!!!")
}

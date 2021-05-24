package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type EpicJSON struct {
	Kenobi   string `json:"kenobi"`
	Grievous string `json:"grievous"`
}

func HandleRequest(port string) {
	http.HandleFunc("/epic", Epic)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln(err)
		return
	}
}

func Epic(w http.ResponseWriter, r *http.Request) {
	_ = r
	var monumental EpicJSON
	monumental.Kenobi = "Hello there"
	monumental.Grievous = "General Kenobi"
	err := json.NewEncoder(w).Encode(monumental)
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Println("Epic received")
}

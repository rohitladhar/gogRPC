package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Inventory struct{
	Id int32 `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
}

type AllInventory struct {
	Data []Inventory `json:"data"`
}
func(app *Config) CreateInventory(w http.ResponseWriter, r *http.Request){

	allInventory := []Inventory{
		{
			Id:1,Name:"Shoes",Description: "Size 10",
		},
		{
			Id:2,Name:"Shirts",Description: "Size 10",
		},
		{
			Id:3,Name:"Skirts",Description: "Size 10",
		},
	}
	
	data := &AllInventory{Data: allInventory}

	out,err := json.Marshal(data)
	if err != nil{
		log.Println(err)
	}

	w.Header().Set("content-type","application/json")
	w.WriteHeader(http.StatusOK)

	_,err = w.Write(out)
	if err != nil{
		log.Println(err)
	}
	
	// w.Header().Set("Content-type","application/json");
	// w.WriteHeader(http.StatusOK)
	// _, err = w.Write(out)
	// if err != nil {
	// 	log.Println(err)
	// }
}
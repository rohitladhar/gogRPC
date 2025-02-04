package main

import (
	"fmt"
	"log"
	"net/http"
)

type Config struct{

}

const PORT = "8080"

func main(){
	log.Println("Starting Broker Service")
	app := &Config{}
	srv := &http.Server{
		Addr: fmt.Sprintf(":%s",PORT),
		Handler: app.routes(),
	}
	log.Println("Starting Server on PORT:",PORT)
	err := srv.ListenAndServe()
	if err != nil{
		log.Println(err)
	}

}
package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT = "8082"
type Config struct{

}

func main(){
	app := &Config{}
	srv := &http.Server{
		Addr: fmt.Sprintf(":%s",PORT),
		Handler: app.routes(),
	}

	log.Println("Server is listening at:",PORT)
	err := srv.ListenAndServe()
	if err != nil{
		log.Println(err)
	}

}
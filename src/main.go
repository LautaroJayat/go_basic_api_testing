package main

import (
	"fmt"
	"log"
	"net/http"
)

func initHandlers() http.Handler{
	r := http.NewServeMux()
	r.HandleFunc("/", sayHiHandler)
	return r
}

func sayHiHandler(w http.ResponseWriter, r *http.Request ){
	w.Write([]byte("Hi, im glad you call\n"))
}

func main(){
	fmt.Println("Hi! We are starting a new server right now!")

	err := http.ListenAndServe(":8080", initHandlers())
	
	if err!= nil{
		log.Fatal("Something went wrong while starting server on port 8080\n")
	}

}
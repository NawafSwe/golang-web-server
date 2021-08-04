package main 

import (
	"net/http"
	"fmt"
	// "time"
	"log"
)
import ( "webserver/pkg/handlers")
// server entry point 
func main() {

	http.HandleFunc("/health", handlers.HandleHealth);
	// starting server 
	fmt.Println("Listening on localhost:6600");
	if err := http.ListenAndServe(":6600", nil); err != nil{
			log.Fatal(err)
	}

}
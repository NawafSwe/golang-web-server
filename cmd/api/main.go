package main 

import (
	"net/http"
	"fmt"
	// "time"
	"log"
)

type Welcome struct{
	Message string 
	Time string
}

// server entry point 
func main() {

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hello!")
	});
	// starting server 
	fmt.Println("Listening on localhost:6600");
	if err := http.ListenAndServe(":6600", nil); err != nil{
			log.Fatal(err)
	}

}
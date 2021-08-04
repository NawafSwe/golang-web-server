package handlers
import(
	"fmt" 
"net/http"
"encoding/json"
)
func HandleHealth(w http.ResponseWriter, r *http.Request)  {
	if r.URL.Path != "/health" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }

    if r.Method != "GET" {
        http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
        return
    }
    fmt.Fprintf(w, "Hello from golang server!")	
}


type Person struct { 
	Age string `json:"age"`
	Name string `json:"name"`
}

func PostHealth(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/healthPost"{
		http.Error(w,"404 not found",http.StatusNotFound)
		return 
	
	}

	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
		return 
       
	}
	var person Person

	if err := json.NewDecoder(r.Body).Decode(&person); err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}
	
w.Header().Set("Content-Type", "application/json")
 personData , err := json.MarshalIndent(person, "", "  "); 
 if err != nil {
	http.Error(w, err.Error(), http.StatusInternalServerError)
	return 
 }
 w.Write(personData)

}
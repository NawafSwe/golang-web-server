package handlers
import(
	"fmt" 
"net/http")
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
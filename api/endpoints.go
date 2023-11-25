package api

import(
	"fmt"
	"net/http"
)

type Endpoints struct{}

func (e *Endpoints)Get(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "GET request received!!")
}
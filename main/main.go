package main

import(
	"net/http"
	"github.com/Shimodaira0910/auth_gogo/Go/api"
)

func main(){
	endPoint := api.Endpoints{}
	http.HandleFunc("/get", endPoint.GET())
	http.ListenAndServe(":8001", nil)
}
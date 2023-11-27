package main

import (
	"fmt"
	"net/http"

	"github.com/Shimodaira0910/auth_gogo/api"
)

func main(){
	fmt.Println("listen start!")
	endPoint := api.Endpoints{}
	http.HandleFunc("/get", endPoint.HandleGetRequest)
	http.HandleFunc("/post", endPoint.HandlePostRequest)

	http.ListenAndServe(":8001", nil)
}

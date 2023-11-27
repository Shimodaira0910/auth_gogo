package api

import (
	"fmt"
	"io"
	"net/http"
	"encoding/json"

	"github.com/Shimodaira0910/auth_gogo/models"
)

type Endpoints struct{}

func (e *Endpoints) HandleGetRequest(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get Request Received!!")
}

func (e *Endpoints) HandlePostRequest(w http.ResponseWriter, r *http.Request){
	if r.Method == "OPTIONS" {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
        w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
        return
	}

	user := models.User{}
	body, err := io.ReadAll(r.Body)
	if err != nil{
		http.Error(w, "リクエストの読み取りエラー", http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &user)
	if err != nil{
		http.Error(w, "リクエストのアンマーシャルエラー", http.StatusInternalServerError)
		return
	}
	
	fmt.Println(body)
	fmt.Println(user.Username)
	w.Write(body)
}


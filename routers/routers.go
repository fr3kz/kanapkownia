package routers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func StartServer(){
	r := mux.NewRouter()
	http.ListenAndServe(":8080",r)
}
package routers

import (
	"../controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func StartServer(){
	r := mux.NewRouter()
	r.HandleFunc("/",controllers.MainPage)
	r.HandleFunc("/create/{kanapka_id}/",controllers.CreateOrder)
	r.HandleFunc("/continue/{zamowienie_id}/",controllers.ContinueOrder)
	r.HandleFunc("/add/{zamowienie_id}/",controllers.Add_to_Order)
	r.HandleFunc("/finish/{zamowienie_id}/",controllers.FinishOrder)
	http.ListenAndServe(":8080",r)
}
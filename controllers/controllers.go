package controllers

import (
	"../models"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"html/template"
	"net/http"
	"strconv"
)

var db,err = gorm.Open("sqlite3","database.db")
func ConnectToDB(){
	if err != nil{
		panic(err.Error())
	}

	defer db.Close()
	//delete this on production
	db.AutoMigrate(&models.Kanapka{})
	db.AutoMigrate(&models.Ocena{})
	db.AutoMigrate(&models.Zamowienie{})
}

func MainPage(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,"main")
}

func CreateOrder(w http.ResponseWriter,r *http.Request){
	vars := mux.Vars(r)
	var kanapka models.Kanapka
	db.Find(&kanapka, "ID = ?",vars["kanapka_id"])

	zamowienie := models.Zamowienie{
		Kanapki:           nil,
		Adres:             "",
		Miasto:            "",
		Kod_pocztowy:      0,
		Cena:              0,
		Czas_dostarczenia: 0,
		Zrobione:          false,
		Dostarczone:       false,
		Oceny:             nil,
	}
	zamowienie.Kanapki = append(zamowienie.Kanapki, kanapka)

	tmp := template.Must(template.ParseFiles("templates/create.html"))
	tmp.Execute(w, struct {
		Zamowienie_id int64
	}{Zamowienie_id:zamowienie.ID})
}

func ContinueOrder(w http.ResponseWriter,r *http.Request){
	vars := mux.Vars(r)
	zamowienie_id,_ := strconv.ParseInt(vars["zamowienie_id"],10,64)

	tmp := template.Must(template.ParseFiles("templates/continue.html"))
	tmp.Execute(w, struct {
		Zamowienie_id int64
	}{Zamowienie_id: zamowienie_id})
}

func Add_to_Order(w http.ResponseWriter,r *http.Request){
	vars := mux.Vars(r)
	zamowienie_id,_ := strconv.ParseInt(vars["zamowienie_id"],10,64)
	tmp := template.Must(template.ParseFiles("templates/continue.html"))
	tmp.Execute(w, struct {
		Zamowienie_id int64
	}{Zamowienie_id: zamowienie_id})
}

func FinishOrder(w http.ResponseWriter,r *http.Request){
	vars := mux.Vars(r)
	zamowienie_id,_ := strconv.ParseInt(vars["zamowienie_id"],10,64)
	tmp := template.Must(template.ParseFiles("templates/continue.html"))
	tmp.Execute(w, struct {
		Zamowienie_id int64
	}{Zamowienie_id: zamowienie_id})

}
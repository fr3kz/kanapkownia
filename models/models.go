package models

import "github.com/jinzhu/gorm"

type Kanapka struct {
	ID int64
	Opis string
	Skladniki string
	Nazwa string
	Cena string
}

type Ocena struct{
	gorm.Model
	Opis string
}

type Zamowienie struct {
	ID int64
	Kanapki []Kanapka
	Adres string
	Miasto string
	Kod_pocztowy int
	Cena float32
	Czas_dostarczenia float32
	Zrobione bool
	Dostarczone bool
	Oceny []Ocena
}

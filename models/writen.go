package models

import (
	db "github.com/JieeiroSst/LapTRWeb/config"
)

type Writen struct {
	ISBN   int `json:"isbn"`
	AuthId int `json:"authId"`
}

func InsertWriten(w Writen) {
	db := db.DbConn()
	insert, err := db.Prepare("insert into Writen(ISBN,AuthId) values(?,?)")
	if err != nil {
		panic(err.Error())
	}
	id := w.ISBN
	authId := w.AuthId
	insert.Exec(id, authId)
	defer db.Close()
}

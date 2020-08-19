package admin

import (
	db "github.com/JieeiroSst/LapTRWeb/config"
)

type Library struct {
	Id   int `json:"id"`
	ISBN int `json:"isbn"`
}

func ShowListLibrary() []Library {
	db := db.DbConn()
	selDB, err := db.Query("select * from Library")
	if err != nil {
		panic(err.Error())
	}

	library := Library{}
	res := []Library{}
	for selDB.Next() {
		var id, isbn int

		err = selDB.Scan(&id, &isbn)
		if err != nil {
			panic(err.Error())
		}

		library.Id = id
		library.ISBN = isbn

		res = append(res, library)
	}
	return res
}

func ShowSingleLibrary(id int) Library {
	db := db.DbConn()
	selDB, err := db.Query("select * from Library where Id=?", id)
	if err != nil {
		panic(err.Error())
	}

	library := Library{}
	for selDB.Next() {
		var id, isbn int

		err = selDB.Scan(&id, &isbn)
		if err != nil {
			panic(err.Error())
		}

		library.Id = id
		library.ISBN = isbn
	}
	return library
}

func InserLibrary(l Library) {
	db := db.DbConn()

	insert, err := db.Prepare("insert into Library(Id,ISBN) values(?,?)")
	if err != nil {
		panic(err.Error())
	}
	id := l.Id
	isbn := l.ISBN
	insert.Exec(id, isbn)
	defer db.Close()
}

func UpdateLibrary(l Library) {
	db := db.DbConn()

	insert, err := db.Prepare("update Library set ISBN=? where Id=?")
	if err != nil {
		panic(err.Error())
	}
	id := l.Id
	isbn := l.ISBN
	insert.Exec(isbn, id)
	defer db.Close()
}

func DeleteLibrary(id int) {
	db := db.DbConn()

	insert, err := db.Prepare("delete from Library where Id=?")
	if err != nil {
		panic(err.Error())
	}
	insert.Exec(id)
	defer db.Close()
}

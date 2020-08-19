package admin

import (
	db "github.com/JieeiroSst/LapTRWeb/config"
)

type Author struct {
	AuthId     int    `json:"auth_id"`
	Name       string `json:"name"`
	Affliation string `json:"affliation"`
	Email      string `json:"email"`
}

func ShowListAuthor() []Author {
	db := db.DbConn()
	seleDB, err := db.Query("select * from Author")
	if err != nil {
		panic(err.Error())
	}

	author := Author{}
	res := []Author{}

	for seleDB.Next() {
		var id int
		var name, afflication, email string

		err = seleDB.Scan(&id, &name, &afflication, &email)
		if err != nil {
			panic(err.Error())
		}
		author.AuthId = id
		author.Name = name
		author.Affliation = afflication
		author.Email = email

		res = append(res, author)
	}
	defer db.Close()
	return res
}

func ShowSingleAuthor(id int) Author {
	db := db.DbConn()
	seleDB, err := db.Query("select * from Author where AuthId=?", id)
	if err != nil {
		panic(err.Error())
	}

	author := Author{}

	for seleDB.Next() {
		var id int
		var name, afflication, email string

		err = seleDB.Scan(&id, &name, &afflication, &email)
		if err != nil {
			panic(err.Error())
		}
		author.AuthId = id
		author.Name = name
		author.Affliation = afflication
		author.Email = email
	}
	defer db.Close()
	return author
}

func InsertAuthor(a Author) {
	db := db.DbConn()
	insert, err := db.Prepare("insert into Author(AuthId,Name,Affliation,Email) values(?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	authID := a.AuthId
	name := a.Name
	affliation := a.Affliation
	email := a.Email
	insert.Exec(authID, name, affliation, email)
	defer db.Close()
}

func UpdateAuthor(a Author) {
	db := db.DbConn()
	update, err := db.Prepare("update Author set Name=? and Affliation=? and Email=? where AuthId=?")
	if err != nil {
		panic(err.Error())
	}

	authID := a.AuthId
	name := a.Name
	affliation := a.Affliation
	email := a.Email
	update.Exec(name, affliation, email, authID)
	defer db.Close()
}

func DeleteAuth(id int) {
	db := db.DbConn()
	delete, err := db.Prepare("delete from Author where AuthId=?")
	if err != nil {
		panic(err.Error())
	}
	delete.Exec(id)
	defer db.Close()
}

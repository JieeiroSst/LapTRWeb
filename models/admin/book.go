package admin

import (
	db "github.com/JieeiroSst/LapTRWeb/config"
)

type Book struct {
	ISBN      int    `json:"json"`
	Title     string `json:"title"`
	Language  string `json:"language"`
	MRP       string `json:"MRP"`
	Publisher string `json:"publisher"`
	PubDate   string `json:"pub_date"`
}

func ShowListBook() []Book {
	db := db.DbConn()
	selDB, err := db.Query("SELECT * FROM Book")
	if err != nil {
		panic(err.Error())
	}
	book := Book{}
	res := []Book{}

	for selDB.Next() {
		var idbn int
		var title, language, mrp, publicsher, pubDate string
		err = selDB.Scan(&idbn, &title, &mrp, &publicsher, &pubDate)
		if err != nil {
			panic(err.Error())
		}
		book.ISBN = idbn
		book.Title = title
		book.Language = language
		book.MRP = mrp
		book.Publisher = publicsher
		book.PubDate = pubDate

		res = append(res, book)
	}
	defer db.Close()
	return res
}

func ShowSingleBook(id int) Book {
	db := db.DbConn()
	selDB, err := db.Query("SELECT * FROM Book where ISBN=?", id)
	if err != nil {
		panic(err.Error())
	}
	book := Book{}

	for selDB.Next() {
		var idbn int
		var title, language, mrp, publicsher, pubDate string
		err = selDB.Scan(&idbn, &title, &mrp, &publicsher, &pubDate)
		if err != nil {
			panic(err.Error())
		}
		book.ISBN = idbn
		book.Title = title
		book.Language = language
		book.MRP = mrp
		book.Publisher = publicsher
		book.PubDate = pubDate
	}
	defer db.Close()
	return book
}

func InsertBook(b Book) {
	db := db.DbConn()
	insert, err := db.Prepare("insert into Book(ISBN,Title,Language,MRP,Publisher,PubDate) values(?,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	isbn := b.ISBN
	title := b.Title
	language := b.Language
	mrp := b.MRP
	publicsher := b.Publisher
	pubdate := b.PubDate
	insert.Exec(isbn, title, language, mrp, publicsher, pubdate)
	defer db.Close()
}

func UpdateBook(b Book) {
	db := db.DbConn()
	update, err := db.Prepare("update Book set Title=? and Language=? and MRP=? and Publisher=? and PubDate=? where ISBN=?")
	if err != nil {
		panic(err.Error())
	}
	isbn := b.ISBN
	title := b.Title
	language := b.Language
	mrp := b.MRP
	publicsher := b.Publisher
	pubdate := b.PubDate
	update.Exec(title, language, mrp, publicsher, pubdate, isbn)
	defer db.Close()
}

func DeleteBook(isbn int) {
	db := db.DbConn()
	delete, err := db.Prepare("delete from Bok where ISBN=?")
	if err != nil {
		panic(err.Error())
	}

	delete.Exec(isbn)
	defer db.Close()
}

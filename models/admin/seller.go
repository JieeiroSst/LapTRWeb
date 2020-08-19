package admin

import (
	db "github.com/JieeiroSst/LapTRWeb/config"
)

type Seller struct {
	ISBN      int `json:"isbn"`
	SellerId  int `json:"sellerid"`
	Type      int `json:"type"`
	NumCopies int `json:"numcopies"`
	Price     int `json:"price"`
}

func ShowListSell() []Seller {
	db := db.DbConn()

	selDb, err := db.Query("select * from Seller")
	if err != nil {
		panic(err.Error())
	}
	seller := Seller{}
	res := []Seller{}
	for selDb.Next() {
		var isbn, sellerID, types, numcopies, price int
		err = selDb.Scan(&isbn, &sellerID, &types, &numcopies, &price)
		if err != nil {
			panic(err.Error())
		}
		seller.ISBN = isbn
		seller.SellerId = sellerID
		seller.Type = types
		seller.NumCopies = numcopies
		seller.Price = price

		res = append(res, seller)
	}
	defer db.Close()
	return res
}

func ShowSigleSell(id int) Seller {
	db := db.DbConn()

	selDb, err := db.Query("select * from Seller where SellerId=?", id)
	if err != nil {
		panic(err.Error())
	}
	seller := Seller{}
	for selDb.Next() {
		var isbn, sellerID, types, numcopies, price int
		err = selDb.Scan(&isbn, &sellerID, &types, &numcopies, &price)
		if err != nil {
			panic(err.Error())
		}
		seller.ISBN = isbn
		seller.SellerId = sellerID
		seller.Type = types
		seller.NumCopies = numcopies
		seller.Price = price
	}
	defer db.Close()
	return seller
}

func InsertSeller(s Seller) {
	db := db.DbConn()
	insert, err := db.Prepare("insert into Seller(ISBN,Type,NumCopies,Price) values(?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	isbn := s.ISBN
	types := s.Type
	numCopies := s.NumCopies
	price := s.Price

	insert.Exec(isbn, types, numCopies, price)

	defer db.Close()
}

func DeleteSeller(id int) {
	db := db.DbConn()
	delete, err := db.Prepare("delete from Seller where SellerId=?")
	if err != nil {
		panic(err)
	}
	delete.Exec(id)
	defer db.Close()
}

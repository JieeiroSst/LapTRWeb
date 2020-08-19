package admin

import (
	db "github.com/JieeiroSst/LapTRWeb/config"
)

type BuySellRecord struct {
	SellerId int `json:"sellerId"`
	BuyerId  int `json:"buyerId"`
	Stype    int `json:"stype"`
	ISBN     int `json:"isbn"`
	Tid      int `json:"tid"`
}

func ShowListBuySell() []BuySellRecord {
	db := db.DbConn()
	selDb, err := db.Query("select * from BuySellRecord")
	if err != nil {
		panic(err.Error())
	}

	buy := BuySellRecord{}
	res := []BuySellRecord{}

	for selDb.Next() {
		var sellerID, buyerId, stype, isbn, tid int

		err = selDb.Scan(&sellerID, &buyerId, &stype, &isbn, &tid)
		if err != nil {
			panic(err.Error())
		}

		buy.SellerId = sellerID
		buy.BuyerId = buyerId
		buy.Stype = stype
		buy.ISBN = isbn
		buy.Tid = tid

		res = append(res, buy)
	}

	defer db.Close()
	return res
}

func ShowSigleBuySell(id int) BuySellRecord {
	db := db.DbConn()
	selDb, err := db.Query("select * from BuySellRecord where SellerID=?", id)
	if err != nil {
		panic(err.Error())
	}

	buy := BuySellRecord{}

	for selDb.Next() {
		var sellerID, buyerId, stype, isbn, tid int

		err = selDb.Scan(&sellerID, &buyerId, &stype, &isbn, &tid)
		if err != nil {
			panic(err.Error())
		}

		buy.SellerId = sellerID
		buy.BuyerId = buyerId
		buy.Stype = stype
		buy.ISBN = isbn
		buy.Tid = tid
	}

	defer db.Close()
	return buy
}

func InsertBuySell(b BuySellRecord) {
	db := db.DbConn()

	insert, err := db.Prepare("insert into BuySellRecord(BuyerId,Stype,ISBN,Tid) values(?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	buyerID := b.BuyerId
	stype := b.Stype
	isbn := b.ISBN
	tid := b.Tid
	insert.Exec(buyerID, stype, isbn, tid)
	defer db.Close()
}

func DeleteBuySell(id int) {
	db := db.DbConn()
	delete, err := db.Prepare("delete * from BuySellRecord where SellerId=?")
	if err != nil {
		panic(err.Error)
	}
	delete.Exec(id)

	defer db.Close()
}

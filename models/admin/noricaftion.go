package admin

import (
	db "github.com/JieeiroSst/LapTRWeb/config"
)

type Notification struct {
	SellerId int `json:"sellerId"`
	BuyerId  int `json:"buyerID"`
	ISBN     int `json:"isbn"`
	Tid      int `json:"tid"`
}

func InsertNotification(n Notification) {
	db := db.DbConn()
	insert, err := db.Prepare("insert into Notification() values(?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	sellerid := n.SellerId
	buyerid := n.BuyerId
	isbn := n.ISBN
	tid := n.Tid
	insert.Exec(sellerid, buyerid, isbn, tid)
}

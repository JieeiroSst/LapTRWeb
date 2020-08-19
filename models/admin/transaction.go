package admin

import (
	"time"

	db "github.com/JieeiroSst/LapTRWeb/config"
)

type Transaction struct {
	Tid           int       `json:"tid"`
	Date          time.Time `json:"date"`
	Total         int       `json:"total"`
	CreditcardNum int       `json:"creditcardNum"`
	CcType        string    `json:"cc_type"`
	CcExpiry      string    `json:"cc_expiry"`
}

func InsertTrasaction(t Transaction) {
	db := db.DbConn()
	insert, err := db.Prepare("insert into Transaction(Tid,Date,Total,CreditcardNum,CcType,CcExpiry) values(?,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	id := t.Tid
	dates := t.Date
	total := t.Total
	CreadNum := t.CreditcardNum
	Cctype := t.CcType
	ccExpiry := t.CcExpiry
	insert.Exec(id, dates, total, CreadNum, Cctype, ccExpiry)
	defer db.Close()
}

func DeleteTrasaction(id int) {
	db := db.DbConn()
	delete, err := db.Prepare("delete from Transaction where Tid=?")
	if err != nil {
		panic(err.Error())
	}
	delete.Exec(id)
	defer db.Close()
}

package admin

import (
	db "github.com/JieeiroSst/LapTRWeb/config"
)

type Recommendation struct {
	ISBN     int `json:"isbn"`
	CourseID int `json:"courseID"`
}

func InsertRecommendation(r Recommendation) {
	db := db.DbConn()
	insert, err := db.Prepare("insert into Recommendation(ISBN,CourseID) values(?,?)")
	if err != nil {
		panic(err.Error())
	}
	isbn := r.ISBN
	courseid := r.CourseID
	insert.Exec(isbn, courseid)
	defer db.Close()
}

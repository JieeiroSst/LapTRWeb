package admin

import (
	db "github.com/JieeiroSst/LapTRWeb/config"
)

type Course struct {
	CourseId   int    `json:"course_id"`
	CourseName string `json:"course name"`
	Department string `json:"department"`
}

func ShowListCourse() []Course {
	db := db.DbConn()
	selDb, err := db.Query("select * from Course")
	if err != nil {
		panic(err.Error())
	}
	course := Course{}
	res := []Course{}
	for selDb.Next() {
		var id int
		var coursename, department string
		err = selDb.Scan(&id, &coursename, &department)
		if err != nil {
			panic(err.Error())
		}

		course.CourseId = id
		course.CourseName = coursename
		course.Department = department

		res = append(res, course)
	}
	defer db.Close()
	return res
}

func ShowsingleCourse(id int) Course {
	db := db.DbConn()
	selDb, err := db.Query("select * from Course where CourseId=?", id)
	if err != nil {
		panic(err.Error())
	}
	course := Course{}
	for selDb.Next() {
		var id int
		var coursename, department string
		err = selDb.Scan(&id, &coursename, &department)
		if err != nil {
			panic(err.Error())
		}

		course.CourseId = id
		course.CourseName = coursename
		course.Department = department
	}
	defer db.Close()
	return course
}

func InsertCourse(c Course) {
	db := db.DbConn()
	insert, err := db.Prepare("insert into Course(CourseId,CourseName,Department) values(?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	id := c.CourseId
	name := c.CourseName
	department := c.Department
	insert.Exec(id, name, department)
	defer db.Close()
}

func UpdateCourse(c Course) {
	db := db.DbConn()
	update, err := db.Prepare("update Course set CourseName=? and Department=? where CourseId=?")
	if err != nil {
		panic(err.Error())
	}
	id := c.CourseId
	name := c.CourseName
	department := c.Department
	update.Exec(name, department, id)
	defer db.Close()
}

func DeleteCourse(id int) {
	db := db.DbConn()
	delete, err := db.Prepare("delete from Course where CourseId=?")
	if err != nil {
		panic(err.Error())
	}

	delete.Exec(id)
	defer db.Close()
}

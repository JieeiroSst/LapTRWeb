package admin

import (
	db "github.com/JieeiroSst/LapTRWeb/config"
)

type Profile struct {
	UserID   int    `json:"userid"`
	UserName string `json:"username"`
	Address  string `json:"address"`
	Password string `json:"password"`
}

func ShowListProfile() []Profile {
	db := db.DbConn()
	selDB, err := db.Query("select * from Profile")
	if err != nil {
		panic(err.Error())
	}

	profile := Profile{}
	res := []Profile{}

	for selDB.Next() {
		var id int
		var username, address, password string

		err = selDB.Scan(&id, &username, &address, &password)
		if err != nil {
			panic(err.Error())
		}
		profile.UserID = id
		profile.UserName = username
		profile.Address = address
		profile.Password = password

		res = append(res, profile)
	}
	defer db.Close()
	return res
}

func ShowSingleProfile(id int) Profile {
	db := db.DbConn()
	selDB, err := db.Query("select * from Profile where UserId=?", id)
	if err != nil {
		panic(err.Error())
	}

	profile := Profile{}

	for selDB.Next() {
		var ids int
		var username, address, password string

		err = selDB.Scan(&ids, &username, &address, &password)
		if err != nil {
			panic(err.Error())
		}
		profile.UserID = ids
		profile.UserName = username
		profile.Address = address
		profile.Password = password
	}
	defer db.Close()
	return profile
}

func UpdateProfile(p Profile) {
	db := db.DbConn()
	update, err := db.Prepare("update Profile set UserName=? and Address=? where UserId=?")
	if err != nil {
		panic(err.Error())
	}

	id := p.UserID
	username := p.UserName
	address := p.Address

	update.Exec(username, address, id)
	defer db.Close()
}

func DeleteProfile(id int) {
	db := db.DbConn()
	delete, err := db.Prepare("delete from Profile where UserId=?")
	if err != nil {
		panic(err.Error())
	}
	delete.Exec(id)
	defer db.Close()
}

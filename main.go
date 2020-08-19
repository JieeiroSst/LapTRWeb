package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JieeiroSst/LapTRWeb/controllers/admin"
	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()

	files := http.FileServer(http.Dir("./public"))
	route.Handle("/", files)

	route.HandleFunc("/admin", admin.HomeAdmin)

	route.HandleFunc("/admin/auth", admin.IndexAuthor)
	route.HandleFunc("/admin/auth/show", admin.IndexSignleauthor)
	route.HandleFunc("/admin/auth/delete", admin.DeleteAuth)
	route.HandleFunc("/admin/auth/edit", admin.EditAuthor)
	route.HandleFunc("/admin/auth/create", admin.CreateAuth)
	route.HandleFunc("/admin/auth/insert", admin.InsertAuth)
	route.HandleFunc("/admin/auth/update", admin.UpdateAuth)

	route.HandleFunc("/admin/book", admin.ShowListBook)
	route.HandleFunc("/admin/book/show", admin.ShowSingleBook)
	route.HandleFunc("/admin/book/delete", admin.DeleteBook)
	route.HandleFunc("/admin/book/edit", admin.EditBook)
	route.HandleFunc("/admin/book/create", admin.CreateBook)
	route.HandleFunc("/admin/book/insert", admin.InsertBook)
	route.HandleFunc("/admin/book/update", admin.UpdateBook)

	route.HandleFunc("/admin/buy-sell", admin.ShowListBuySell)
	route.HandleFunc("/admin/buy-sell/show", admin.ShowSingleSell)
	route.HandleFunc("/admin/buy-sell/delete", admin.DeleteSell)
	route.HandleFunc("/admin/buy-sell/edit", admin.EditBuySell)
	route.HandleFunc("/admin/buy-sell/create", admin.CreateBuysell)
	route.HandleFunc("/admin/buy-sell/insert", admin.InsertBuySell)

	route.HandleFunc("/admin/course", admin.ShowListcourse)
	route.HandleFunc("/admin/course/show", admin.ShowSigleCourse)
	route.HandleFunc("/admin/course/delete", admin.DeleteCourse)
	route.HandleFunc("/admin/course/create", admin.CreateCourse)
	route.HandleFunc("/admin/course/insert", admin.InsertCourse)
	route.HandleFunc("/admin/course/update", admin.UpdateCourse)

	route.HandleFunc("/admin/library", admin.ShowListLibrary)
	route.HandleFunc("/admin/library/show", admin.ShowSingleLibrary)
	route.HandleFunc("/admin/library/delete", admin.DeleteLibrary)
	route.HandleFunc("/admin/library/edit", admin.EditLibrary)
	route.HandleFunc("/admin/library/create", admin.CreateLibrary)
	route.HandleFunc("/admin/library/update", admin.UpdateLibrary)

	route.HandleFunc("/admin/profile", admin.ShowListProfile)
	route.HandleFunc("/admin/profile/show", admin.ShowSingleProfile)
	route.HandleFunc("/admin/profile/delete", admin.DeleteProfile)
	route.HandleFunc("/admin/profile/edit", admin.EditProfile)
	route.HandleFunc("/admin/profile/update", admin.UpdateProdile)

	route.HandleFunc("/admin/selle", admin.ShowListSeller)
	route.HandleFunc("/admin/selle/show", admin.ShowSingleSell)
	route.HandleFunc("/admin/selle/delete", admin.DeleteSell)

	fmt.Println("server running port 9000...")
	log.Fatal(http.ListenAndServe(":9000", route))
}

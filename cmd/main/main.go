package main

import (
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/turuu0710/go-mysql/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes()
	http.Handle("/", r)
	logg.Fatal(http.ListenAndServe("localhost:9010", r))
}

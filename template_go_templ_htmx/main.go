package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/exp/slog"

	components "template/client"
	"template/db"
	"template/services"

	"github.com/a-h/templ"
)

func main() {
	fmt.Println("Main")

	logg := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	db := db.NewDBInstance(os.Getenv("DB_NAME"))
	services.GlobalService(logg, db)

	component := components.Index()
	http.Handle("/", templ.Handler(component))

	fmt.Println("On http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}

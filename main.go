package main

import (
	"net/http"

	"github.com/door-service/api/controllers"
)

func main() {
	mux := http.NewServeMux()

	QC := controllers.NewQuoteController()
	mux.HandleFunc("/quotes/", QC.GetAll)

	http.ListenAndServe(":8080", mux)
}

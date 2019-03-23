package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"

	"./common"
	"./model"
)

var router = mux.NewRouter()

func main() {
	fmt.Println("GOSHRT (C) 2019 Nathan")

	common.AddUrl(model.Url{To: "https://google.com", Description: "Google"})
	common.AddUrl(model.Url{To: "https://example.com", Description: "Example"})
	common.AddUrl(model.Url{To: "https://www.icann.org", Description: "Icann"})

	router.HandleFunc("/get-token", common.GetToken).Methods("GET")

	router.Handle("/redirects", common.IsAuthorized(common.GetRedirectInfos)).Methods("GET")
	router.Handle("/createredirect", common.IsAuthorized(common.CreateRedirect)).Methods("POST")
	router.HandleFunc("/{id}", common.GetRedirect).Methods("GET")
	router.HandleFunc("/{id}/info", common.GetRedirectInfo).Methods("GET")

	http.Handle("/", router)
	_ = http.ListenAndServe(":8080", nil)
}

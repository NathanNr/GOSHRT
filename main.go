package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"./common"
	"./model"
)

var router = mux.NewRouter()

func main() {
	fmt.Println("GOSHRT (C) 2019 Nathan")

	common.SetUrls(append(common.GetUrls(), model.Url{ID: "1", To: "https://google.com", Description: "Google"}))
	common.SetUrls(append(common.GetUrls(), model.Url{ID: "2", To: "https://example.com", Description: "Example"}))
	common.SetUrls(append(common.GetUrls(), model.Url{ID: "3", To: "https://www.icann.org", Description: "Icann"}))

	router.HandleFunc("/urls", common.GetRedirectInfos).Methods("GET")
	router.HandleFunc("/{id}", common.GetRedirect).Methods("GET")
	router.HandleFunc("/{id}/info", common.GetRedirectInfo).Methods("GET")

	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}

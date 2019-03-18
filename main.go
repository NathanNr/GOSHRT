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
	fmt.Println("GOSHRT by Nathan")

	common.SetUrls(append(common.GetUrls(), model.Url{ID: "1", To: "https://google.com", Description: "Google"}))

	router.HandleFunc("/{id}", common.GetRedirect).Methods("GET")
	router.HandleFunc("/{id}/info", common.GetRedirectInfo).Methods("GET")
	router.HandleFunc("/urls", common.GetRedirectInfos).Methods("GET")

	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}

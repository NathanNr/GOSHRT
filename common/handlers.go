package common

import (
	"../model"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

var urls []model.Url

func GetUrls() []model.Url {
	return urls
}

func SetUrls (urlss []model.Url) {
	urls = urlss
}

func GetRedirect(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range urls {
		if item.ID == params["id"] {
			http.Redirect(w, r, item.To, 301)
		}
	}
}

func CreateRedirect(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST create redirect")
	params := mux.Vars(r)
	var url model.Url
	_ = json.NewDecoder(r.Body).Decode(&url)
	url.ID = params["id"]
	url.Creationtime = time.Now().Unix()
	urls = append(urls, url)
	json.NewEncoder(w).Encode(url)
}

func GetRedirectInfo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range urls {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&model.Url{})
}

func GetRedirectInfos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(urls)
}

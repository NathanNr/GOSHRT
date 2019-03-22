package common

import (
	"../model"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

var urlCount = 0
var urls []model.Url

func AddUrl (urlToAdd model.Url) {
	urlCount = urlCount + 1
	urlToAdd.ID = strconv.Itoa(urlCount)
	urlToAdd.Creationtime = time.Now().Unix()
	urls = append(urls, urlToAdd)
}

func GetRedirect(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range urls {
		if item.ID == params["id"] {
			http.Redirect(w, r, item.To, 301)
		}
	}
	http.NotFound(w, r)
}

func CreateRedirect(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var url model.Url
	_ = json.NewDecoder(r.Body).Decode(&url)
	AddUrl(url)
	_ = json.NewEncoder(w).Encode(url)
}

func GetRedirectInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range urls {
		if item.ID == params["id"] {
			_ = json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.NotFound(w, r)
}

func GetRedirectInfos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(urls)
}

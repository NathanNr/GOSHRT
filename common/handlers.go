package common

import (
	"../model"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

var urls []model.Url

func GetRedirect(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range urls {
		if item.ID == params["id"] {
			http.Redirect(w, r, item.To, 301)
		}
	}
}

func GetUrls() []model.Url {
	return urls
}

func SetUrls (urlss []model.Url) {
	urls = urlss
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

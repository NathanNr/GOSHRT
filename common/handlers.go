package common

import (
	"../model"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
)

var urlCount = 0
var urls []model.Url

func AddUrl(urlToAdd model.Url) {
	urlCount = urlCount + 1
	urlToAdd.ID = strconv.Itoa(urlCount)
	urlToAdd.Creationtime = time.Now().Unix()
	urls = append(urls, urlToAdd)
}

func RemoveUrl(urlIdToRemove string) {
	for i, item := range urls {
		if item.ID == urlIdToRemove {
			urls[i] = model.Url{}
		}
	}
}

func GetRedirect(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "private, max-age=90")
	params := mux.Vars(r)
	for _, item := range urls {
		if item.ID == params["id"] {
			http.Redirect(w, r, item.To, 301)
			return
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
	w.WriteHeader(http.StatusCreated)
}

func GetRedirectInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range urls {
		if item.ID == params["id"] {
			err := json.NewEncoder(w).Encode(item)
			if err != nil {
				log.Fatal(err)
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		}
	}
	http.NotFound(w, r)
}

func GetRedirectInfos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(urls)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func DeleteRedirect(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	RemoveUrl(params["id"])
	w.WriteHeader(http.StatusNoContent)
}

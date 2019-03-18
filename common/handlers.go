package common

import (
	"../model"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var urls []model.Url

func GetRedirect(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var body = ""
	for _, item := range urls {
		if item.ID == params["id"] {
			body = "<html><head>" +
				"<title>" + item.Description + "</title>" +
				"<meta http-equiv=\"refresh\" content=\"0; URL=" + item.To + "\">" +
				"</head><body>" +
				"<p>Redirecting to <a href=\"" + item.To + "\">" + item.Description + "</a>...</p>" +
				"</body></html>"
		}
	}
	fmt.Fprintf(w, body)

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

package common

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"strings"
	"time"
)

var fmtSecret = []byte("secret") // TODO

func GetToken(w http.ResponseWriter, r *http.Request) {
	validToken, err := GenerateJWT("get-token", 60)
	if err != nil {
		fmt.Println("Failed to generate token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = fmt.Fprintf(w, validToken)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func GenerateJWT(client string, exp time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = client
	claims["exp"] = time.Now().Add(time.Minute * exp).Unix()

	tokenString, _ := token.SignedString(fmtSecret)

	return tokenString, nil
}

func IsAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Authorization"] != nil {
			var auth = r.Header["Authorization"][0]
			auth = strings.Replace(auth, "Bearer ", "", 1)
			token, err := jwt.Parse(auth, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("error")
				}
				return fmtSecret, nil
			})

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			if token.Valid {
				endpoint(w, r)
			} else {
				w.WriteHeader(http.StatusForbidden)
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	})
}

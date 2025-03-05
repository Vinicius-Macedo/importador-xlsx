package helpers

import (
	"log"
	"net/http"
)

func GetCookie(r *http.Request, name string) (*http.Cookie, error) {
	cookie, err := r.Cookie(name)
	if err != nil {
		log.Printf("Cookie %s not found\n", name)
		return nil, err
	}

	return cookie, nil
}

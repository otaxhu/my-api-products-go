package main

import (
	"log"
	"net/http"
	"time"

	"github.com/otaxhu/serverX"
)

func Logging() serverX.Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				log.Println(r.URL.Path, time.Since(start))
			}()
			f(w, r)
		}
	}
}

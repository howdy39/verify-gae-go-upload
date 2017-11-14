package app

import (
	"net/http"
	"verify-gae-go-upload/gae"
)

func init() {
	http.HandleFunc("/", gae.IndexHandler)
	http.HandleFunc("/serve/", gae.ServeHandler)
	http.HandleFunc("/upload", gae.UploadHandler)
}
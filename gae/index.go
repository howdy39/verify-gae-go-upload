package gae

import (
	"net/http"
	"google.golang.org/appengine/blobstore"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"html/template"
	"context"
)

var (
	rootTemplate = template.Must(template.ParseFiles("templates/index.html"))
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	uploadURL, err := blobstore.UploadURL(ctx, "/upload", nil)
	if err != nil {
		serveError(ctx, err)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	err = rootTemplate.Execute(w, uploadURL)
	if err != nil {
		serveError(ctx, err)
	}
}

func serveError(ctx context.Context, err error) {
	log.Errorf(ctx, "%v", err)
}
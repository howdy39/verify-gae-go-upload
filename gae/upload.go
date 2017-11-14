package gae

import (
	"net/http"
	"google.golang.org/appengine/blobstore"
	"google.golang.org/appengine"
	"errors"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	blobs, _, err := blobstore.ParseUpload(r)
	if err != nil {
		serveError(ctx, err)
		return
	}

	file := blobs["file"]
	if len(file) == 0 {
		serveError(ctx, errors.New("no file uploaded"))
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	http.Redirect(w, r, "/serve/?blobKey="+string(file[0].BlobKey), http.StatusFound)
}

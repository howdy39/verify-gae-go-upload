package gae

import (
	"net/http"
	"google.golang.org/appengine/blobstore"
	"google.golang.org/appengine"
)

func ServeHandler(w http.ResponseWriter, r *http.Request) {
	blobstore.Send(w, appengine.BlobKey(r.FormValue("blobKey")))
}

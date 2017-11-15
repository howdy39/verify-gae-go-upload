package gae

import (
	"net/http"
	"google.golang.org/appengine/blobstore"
	"google.golang.org/appengine"
	"encoding/csv"
	"strings"
	"fmt"
)

func ServeHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	//blobstore.Send(w, appengine.BlobKey(r.FormValue("blobKey")))

	out := blobstore.NewReader(ctx, appengine.BlobKey(r.FormValue("blobKey")))
	data := csv.NewReader(out)

	// csv全出力
	cols, err := data.Read()
	// 最初の読み込みでエラーが出るなら、Blobは存在しないってよ
	if err != nil {
		http.Error(w, "blob file doesn't exist", http.StatusNotFound)
		return
	}
	for err == nil {
		fmt.Print(cols[0])
		w.Write([]byte(cols[0] + "    "))
		w.Write([]byte(strings.Join(cols, ",")))
		w.Write([]byte("\n"))
		cols, err = data.Read()
	}
}

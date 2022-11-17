package service

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	application "tempest-data-service/pkg/application/entities"

	"cloud.google.com/go/storage"
	"github.com/gorilla/mux"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func newDataOperation(r *mux.Router) {
	r.HandleFunc("/data/test/{username}", tt).Methods(http.MethodGet)
}

func tt(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	username := params["username"]

	client, err := storage.NewClient(context.Background(), option.WithoutAuthentication())
	if err != nil {
		body := application.NewResponse(nil, fmt.Errorf("error"))
		writeReponse(w, r, body)
		return
	}
	bkt := client.Bucket("test-bucket-gfdjgfdhg")

	var names []string
	it := bkt.Objects(context.Background(), &storage.Query{
		Prefix: username,
	})
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		names = append(names, attrs.Name)
	}

	rdr, _ := bkt.Object(names[2]).NewReader(context.Background())

	byt, _ := io.ReadAll(rdr)

	body := application.NewResponse(string(byt))

	writeReponse(w, r, body)
}

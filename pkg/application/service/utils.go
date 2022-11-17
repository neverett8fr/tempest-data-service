package service

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"tempest-data-service/pkg/config"
	st "tempest-data-service/pkg/infra/storage"

	"github.com/gorilla/mux"
)

var (
	StorageProvider st.StorageProvider
)

const (
	username = "username"
	item     = "item"
)

func NewServiceRoutes(r *mux.Router, conf config.Config) {
	sp, err := st.InitialiseStorageProvider(
		context.Background(),
		conf.Storage.BucketName,
	)
	if err != nil {
		log.Printf("error initialising storage provider, err %v", err)
	}

	StorageProvider = sp

	newDataInformation(r)
	newDataOperation(r)
}

func writeReponse(w http.ResponseWriter, body interface{}) {

	reponseBody, err := json.Marshal(body)
	if err != nil {
		log.Printf("error converting reponse to bytes, err %v", err)
	}
	w.Header().Add("Content-Type", "application/json")

	_, err = w.Write(reponseBody)
	if err != nil {
		log.Printf("error writing response, err %v", err)
	}
}

package service

import (
	"fmt"
	"net/http"
	application "tempest-data-service/pkg/application/entities"

	"github.com/gorilla/mux"
)

func newDataInformation(r *mux.Router) {
	r.HandleFunc("/test/{text}", testHandler).Methods(http.MethodGet)
	r.HandleFunc("/data/{username}", userFileNames).Methods(http.MethodGet)

}

func userFileNames(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	username := params[username]

	fileNames, err := StorageProvider.GetFileInformation(r.Context(), username)
	if err != nil {
		body := application.NewResponse(nil, err)
		writeReponse(w, r, body)
		return
	}

	body := application.NewResponse(fileNames)
	writeReponse(w, r, body)
}

func testHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	text := params["text"]

	body := application.NewResponse(fmt.Sprintf("test: %v", text))

	writeReponse(w, r, body)
}

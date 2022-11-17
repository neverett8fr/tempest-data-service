package service

import (
	"net/http"
	application "tempest-data-service/pkg/application/entities"

	"github.com/gorilla/mux"
)

func newDataOperation(r *mux.Router) {
	r.HandleFunc("/data/{username}/{item}", userFileContent).Methods(http.MethodGet)
}

func userFileContent(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	username := params[username]
	item := params[item]

	fileContent, err := StorageProvider.GetFileContent(
		r.Context(), username, item,
	)
	if err != nil {
		body := application.NewResponse(nil, err)
		writeReponse(w, r, body)
		return
	}

	body := application.NewResponse(string(fileContent))
	writeReponse(w, r, body)

}

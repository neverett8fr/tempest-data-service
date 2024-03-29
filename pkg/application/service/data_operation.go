package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	application "tempest-data-service/pkg/application/entities"

	"github.com/gorilla/mux"
)

func newDataOperation(r *mux.Router) {
	r.HandleFunc("/data/{username}/{item}", userFileDownloadSmall).Methods(http.MethodGet)
	r.HandleFunc("/data/{username}", userFileUploadSmall).Methods(http.MethodPost)
}

func userFileDownloadSmall(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	usr := params[username]
	it := params[item]

	fileContent, err := StorageProvider.GetFileContent(
		r.Context(), usr, it,
	)
	if err != nil {
		body := application.NewResponse(nil, err)
		writeReponse(w, body)
		return
	}

	body := application.NewResponse(string(fileContent))
	writeReponse(w, body)

}

func userFileUploadSmall(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	usr := params[username]

	bodyIn, err := ioutil.ReadAll(r.Body)
	if err != nil {
		body := application.NewResponse(nil, err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		writeReponse(w, body)
		return
	}
	fileData := application.FileData{}
	err = json.Unmarshal(bodyIn, &fileData)
	if err != nil {
		body := application.NewResponse(nil, err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		writeReponse(w, body)
		return
	}

	err = StorageProvider.UploadSmallFile(
		r.Context(),
		usr,
		fileData.FileName,
		[]byte(fileData.FileContent),
	)
	if err != nil {
		body := application.NewResponse(nil, err)
		writeReponse(w, body)
		return
	}

	body := application.NewResponse("File successfully uploaded")
	writeReponse(w, body)

}

package documents

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Mickdevv/savefuel-backend/api"
)

func getDocuments(serverCfg *api.ServerConfig, w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("JWT SECRET :" + serverCfg.JWT_SECRET))
}

func uploadDocument(serverCfg *api.ServerConfig, w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(200000000)
	if err != nil {
		api.RespondWithError(w, http.StatusInternalServerError, "File too large!", err)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		api.RespondWithError(w, http.StatusInternalServerError, "Error with file in request", err)
		return
	}

	defer file.Close()

	fmt.Println("File info : "+handler.Filename, handler.Size, handler.Header)

	filename := filepath.Base(handler.Filename)

	destinationPath := filepath.Join("static", "documents", filename)

	dst, err := os.Create(destinationPath)
	if err != nil {
		api.RespondWithError(w, http.StatusInternalServerError, "Error saving file", err)
	}

	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		api.RespondWithError(w, http.StatusInternalServerError, "Error writing file", err)
	}
	type response struct {
		Message string `json:"message"`
	}

	api.RespondWithJSON(w, http.StatusOK, response{Message: "File successfully uploaded"})

}

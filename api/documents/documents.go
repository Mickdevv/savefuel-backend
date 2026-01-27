package documents

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Mickdevv/savefuel-backend/api"
	"github.com/Mickdevv/savefuel-backend/internal/database"
	"github.com/google/uuid"
)

func getDocuments(serverCfg *api.ServerConfig, w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("JWT SECRET :" + serverCfg.JWT_SECRET))
}

func uploadDocument(serverCfg *api.ServerConfig, w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Title       string `json:"title"`
		Filename    string `json:"filename"`
		Locale      string `json:"locale"`
		Description string `json:"description"`
		Priority    int32  `json:"priority"`
	}

	metadata := r.FormValue("metadata")

	params := parameters{}
	err := json.Unmarshal([]byte(metadata), &params)
	// decoder := json.NewDecoder(metadata)
	// err = decoder.Decode(&params)
	if err != nil {
		api.RespondWithError(w, http.StatusBadRequest, "Error decoding metadata", err)
		return
	}

	err = r.ParseMultipartForm(10 >> 20)
	if err != nil {
		api.RespondWithError(w, http.StatusInternalServerError, "File upload error", err)
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

	destinationPath := filepath.Join(serverCfg.STATIC_FILES_DIR, "documents", filename)
	fileTypeSlice := strings.Split(filename, ".")
	fileType := fileTypeSlice[len(fileTypeSlice)-1]

	dst, err := os.Create(destinationPath)
	if err != nil {
		api.RespondWithError(w, http.StatusInternalServerError, "Error saving file", err)
		return
	}

	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		api.RespondWithError(w, http.StatusInternalServerError, "Error writing file", err)
		return
	}

	document, err := serverCfg.DB.CreateDocument(r.Context(), database.CreateDocumentParams{
		Title:       params.Title,
		Filename:    params.Filename,
		Locale:      params.Locale,
		Description: params.Description,
		Priority:    params.Priority,
		Path:        destinationPath,
		Filetype:    fileType,
	})
	if err != nil {
		api.RespondWithError(w, http.StatusInternalServerError, "Error adding document to database", err)
		return
	}

	type response struct {
		Message string `json:"message"`
		Data    struct {
			ID          uuid.UUID `json:"id"`
			CreatedAt   time.Time `json:"created_at"`
			Title       string    `json:"title"`
			Filename    string    `json:"filename"`
			Locale      string    `json:"locale"`
			Description string    `json:"description"`
			Priority    int32     `json:"priority"`
			Path        string    `json:"path"`
			Filetype    string    `json:"filetype"`
		} `json:"data"`
	}
	var res response

	res.Message = "File successfully uploaded"
	res.Data.ID = document.ID
	res.Data.CreatedAt = document.CreatedAt
	res.Data.Title = document.Title
	res.Data.Filename = document.Filename
	res.Data.Locale = document.Locale
	res.Data.Description = document.Description
	res.Data.Priority = document.Priority
	res.Data.Path = destinationPath
	res.Data.Filetype = fileType

	api.RespondWithJSON(w, http.StatusOK, res)
}

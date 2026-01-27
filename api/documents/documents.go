package documents

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/Mickdevv/savefuel-backend/api"
	"github.com/Mickdevv/savefuel-backend/internal/database"
	"github.com/google/uuid"
)

type document struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	Title       string    `json:"title"`
	Filename    string    `json:"filename"`
	Locale      string    `json:"locale"`
	Description string    `json:"description"`
	Priority    int32     `json:"priority"`
	Path        string    `json:"path"`
	Filetype    string    `json:"filetype"`
	Visible     bool      `json:"visible"`
}

func updateDocument(serverCfg *api.ServerConfig, w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	type parameters struct {
		Title       string `json:"title"`
		Locale      string `json:"locale"`
		Description string `json:"description"`
		Priority    int32  `json:"priority"`
		Visible     bool   `json:"visible"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		api.RespondWithError(w, http.StatusBadRequest, "Could not decode json payload", err)
		return
	}

	idStr := r.PathValue("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		api.RespondWithError(w, http.StatusBadRequest, "Invalid document ID", err)
		return
	}

	d, err := serverCfg.DB.UpdateDocument(r.Context(), database.UpdateDocumentParams{
		Title:       params.Title,
		ID:          id,
		Description: params.Description,
		Priority:    params.Priority,
		Visible:     params.Visible,
		Locale:      params.Locale,
	})
	if err != nil {
		api.RespondWithError(w, http.StatusInternalServerError, "Error updating database record", err)
		return
	}

	type response struct {
		Data document `json:"data"`
	}

	res := response{Data: document{
		ID:          d.ID,
		CreatedAt:   d.CreatedAt,
		Title:       d.Title,
		Filename:    d.Filename,
		Locale:      d.Locale,
		Description: d.Description,
		Priority:    d.Priority,
		Path:        d.Path,
		Filetype:    d.Filetype,
		Visible:     d.Visible,
	}}

	api.RespondWithJSON(w, http.StatusOK, res)
}

func deleteDocument(serverCfg *api.ServerConfig, w http.ResponseWriter, r *http.Request) {
	type response struct {
		Data uuid.UUID `json:"data"`
	}

	idstr := r.PathValue("id")
	id, err := uuid.Parse(idstr)
	if err != nil {
		api.RespondWithError(w, http.StatusInternalServerError, "Invalid id", err)
		return
	}

	d, err := serverCfg.DB.GetDocument(r.Context(), id)
	if err != nil {
		api.RespondWithError(w, http.StatusInternalServerError, "Error fetching database records", err)
		return
	}

	err = os.Remove(path.Join(serverCfg.STATIC_FILES_DIR, "documents", d.Filename))
	if err != nil {
		api.RespondWithError(w, http.StatusInternalServerError, "Error deleting file", err)
		return
	}

	err = serverCfg.DB.DeleteDocument(r.Context(), id)
	if err != nil {
		api.RespondWithError(w, http.StatusInternalServerError, "Error deleting the document from the database", err)
		return
	}

	api.RespondWithJSON(w, http.StatusOK, response{Data: id})
}

func getDocumentById(serverCfg *api.ServerConfig, w http.ResponseWriter, r *http.Request) {
	type response struct {
		Data document `json:"data"`
	}

	idStr := r.PathValue("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		api.RespondWithError(w, http.StatusBadRequest, "Invalid document ID", err)
		return
	}

	d, err := serverCfg.DB.GetDocument(r.Context(), id)
	if err != nil {
		api.RespondWithError(w, http.StatusInternalServerError, "Error fetching database records", err)
		return
	}

	res := response{Data: document{
		ID:          d.ID,
		CreatedAt:   d.CreatedAt,
		Title:       d.Title,
		Filename:    d.Filename,
		Locale:      d.Locale,
		Description: d.Description,
		Priority:    d.Priority,
		Path:        d.Path,
		Filetype:    d.Filetype,
	}}

	api.RespondWithJSON(w, http.StatusOK, res)
}

func getDocuments(serverCfg *api.ServerConfig, w http.ResponseWriter, r *http.Request) {
	type response struct {
		Data []document `json:"data"`
	}
	res := response{Data: []document{}}

	documents, err := serverCfg.DB.GetDocuments(r.Context())
	if err != nil {
		api.RespondWithError(w, http.StatusInternalServerError, "Error fetching database records", err)
		return
	}

	for _, d := range documents {
		res.Data = append(res.Data, document{
			ID:          d.ID,
			CreatedAt:   d.CreatedAt,
			Title:       d.Title,
			Filename:    d.Filename,
			Locale:      d.Locale,
			Description: d.Description,
			Priority:    d.Priority,
			Path:        d.Path,
			Filetype:    d.Filetype,
		})

	}

	api.RespondWithJSON(w, http.StatusOK, res)
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

	uploaded_document, err := serverCfg.DB.CreateDocument(r.Context(), database.CreateDocumentParams{
		Title:       params.Title,
		Filename:    filename,
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
		Message string   `json:"message"`
		Data    document `json:"data"`
	}
	var res response

	res.Message = "File successfully uploaded"
	res.Data.ID = uploaded_document.ID
	res.Data.CreatedAt = uploaded_document.CreatedAt
	res.Data.Title = uploaded_document.Title
	res.Data.Filename = uploaded_document.Filename
	res.Data.Locale = uploaded_document.Locale
	res.Data.Description = uploaded_document.Description
	res.Data.Priority = uploaded_document.Priority
	res.Data.Path = destinationPath
	res.Data.Filetype = fileType

	api.RespondWithJSON(w, http.StatusOK, res)
}

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

	"github.com/Mickdevv/savefuel-backend/api"
	"github.com/Mickdevv/savefuel-backend/internal/database"
	"github.com/google/uuid"
)

func updateDocument(serverCfg *api.ServerConfig, w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	params := DocumentPayload{}
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
		Active:      params.Active,
		Locale:      params.Locale,
	})
	if err != nil {
		api.RespondWithError(w, http.StatusInternalServerError, "Error updating database record", err)
		return
	}

	type response struct {
		Data Document `json:"data"`
	}

	res := response{Data: Document{
		ID:          d.ID,
		CreatedAt:   d.CreatedAt,
		Title:       d.Title,
		Filename:    d.Filename,
		Locale:      d.Locale,
		CategoryID:  d.CategoryID,
		Description: d.Description,
		Priority:    d.Priority,
		Path:        d.Path,
		Filetype:    d.Filetype,
		Active:      d.Active,
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
		Data Document `json:"data"`
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

	res := response{Data: Document{
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
		Data []Document `json:"data"`
	}
	res := response{Data: []Document{}}

	documents, err := serverCfg.DB.GetDocuments(r.Context())
	if err != nil {
		api.RespondWithError(w, http.StatusInternalServerError, "Error fetching database records", err)
		return
	}

	for _, d := range documents {
		res.Data = append(res.Data, Document{
			ID:          d.ID,
			CreatedAt:   d.CreatedAt,
			Title:       d.Title,
			Filename:    d.Filename,
			Locale:      d.Locale,
			Description: d.Description,
			Priority:    d.Priority,
			Path:        d.Path,
			Filetype:    d.Filetype,
			CategoryID:  d.CategoryID,
		})

	}

	api.RespondWithJSON(w, http.StatusOK, res)
}

func uploadDocument(serverCfg *api.ServerConfig, w http.ResponseWriter, r *http.Request) {

	metadata := r.FormValue("metadata")

	params := UploadDocumentPayload{}
	err := json.Unmarshal([]byte(metadata), &params)
	if err != nil {
		fmt.Println(metadata, err)
		// api.RespondWithError(w, http.StatusBadRequest, "Error decoding metadata", err)
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
		CategoryID:  params.CategoryID,
	})
	if err != nil {
		api.RespondWithError(w, http.StatusInternalServerError, "Error adding document to database", err)
		return
	}

	type response struct {
		Message string   `json:"message"`
		Data    Document `json:"data"`
	}
	var res response

	res.Message = "File successfully uploaded"
	res.Data.ID = uploaded_document.ID
	res.Data.CreatedAt = uploaded_document.CreatedAt
	res.Data.Title = uploaded_document.Title
	res.Data.Filename = uploaded_document.Filename
	res.Data.Locale = uploaded_document.Locale
	res.Data.CategoryID = uploaded_document.CategoryID
	res.Data.Description = uploaded_document.Description
	res.Data.Priority = uploaded_document.Priority
	res.Data.Path = destinationPath
	res.Data.Filetype = fileType

	api.RespondWithJSON(w, http.StatusOK, res)
}

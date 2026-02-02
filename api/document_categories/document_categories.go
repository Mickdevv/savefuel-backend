package document_categories

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Mickdevv/savefuel-backend/api"
	"github.com/Mickdevv/savefuel-backend/internal/database"
	"github.com/google/uuid"
)

func getDocuments(serverCfg *api.ServerConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		doc_cats, err := serverCfg.DB.GetDocumentCategories(r.Context())
		if err != nil {
			api.RespondWithError(w, http.StatusNotFound, "Not found", err)
			return
		}

		type document struct {
			ID        uuid.UUID `json:"id"`
			Name      string    `json:"name"`
			Active    bool      `json:"active"`
			CreatedAt time.Time `json:"created_at"`
			UpdatedAt time.Time `json:"updated_at"`
		}

		var doc_cats_from_database []document
		for _, doc_cat := range doc_cats {
			doc_cats_from_database = append(doc_cats_from_database, document{
				ID:        doc_cat.ID,
				Name:      doc_cat.Name,
				Active:    doc_cat.Active,
				CreatedAt: doc_cat.CreatedAt,
				UpdatedAt: doc_cat.UpdatedAt,
			})
		}

		api.RespondWithJSON(w, http.StatusOK, doc_cats_from_database)
	}

}
func getDocumentById(serverCfg *api.ServerConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			api.RespondWithError(w, http.StatusBadRequest, "Invalid ID format", err)
			return
		}

		doc_cat, err := serverCfg.DB.GetDocumentCategory(r.Context(), id)
		if err != nil {
			api.RespondWithError(w, http.StatusNotFound, "ID not found", err)
			return
		}

		type response struct {
			ID        uuid.UUID `json:"id"`
			Name      string    `json:"name"`
			Active    bool      `json:"active"`
			CreatedAt time.Time `json:"created_at"`
			UpdatedAt time.Time `json:"updated_at"`
		}

		api.RespondWithJSON(w, http.StatusOK, response{
			ID:        doc_cat.ID,
			Name:      doc_cat.Name,
			Active:    doc_cat.Active,
			CreatedAt: doc_cat.CreatedAt,
			UpdatedAt: doc_cat.UpdatedAt,
		})
	}

}

func deleteDocumentCategory(serverCfg *api.ServerConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			api.RespondWithError(w, http.StatusBadRequest, "Invalid ID format", err)
			return
		}

		err = serverCfg.DB.DeleteDocumentCategory(r.Context(), id)
		if err != nil {
			api.RespondWithError(w, http.StatusNotFound, "ID not found", err)
			return
		}

		w.WriteHeader(http.StatusOK)

	}
}

func updateDocumentCategory(serverCfg *api.ServerConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		idStr := r.PathValue("id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			api.RespondWithError(w, http.StatusBadRequest, "Invalid ID format", err)
			return
		}

		type parameters struct {
			Name   string `json:"name"`
			Active bool   `json:"active"`
		}
		params := parameters{}

		defer r.Body.Close()
		decoder := json.NewDecoder(r.Body)
		err = decoder.Decode(&params)
		if err != nil {
			api.RespondWithError(w, http.StatusBadRequest, "Payload error", err)
			return
		}

		doc_cat, err := serverCfg.DB.UpdateDocumentCategory(r.Context(), database.UpdateDocumentCategoryParams{ID: id, Name: params.Name, Active: params.Active})
		if err != nil {
			api.RespondWithError(w, http.StatusBadRequest, "Error updating in database", err)
			return
		}

		type response struct {
			ID        uuid.UUID `json:"id"`
			Name      string    `json:"name"`
			Active    bool      `json:"active"`
			CreatedAt time.Time `json:"created_at"`
			UpdatedAt time.Time `json:"updated_at"`
		}

		res := response{
			ID:        doc_cat.ID,
			Name:      doc_cat.Name,
			Active:    doc_cat.Active,
			CreatedAt: doc_cat.CreatedAt,
			UpdatedAt: doc_cat.UpdatedAt,
		}
		api.RespondWithJSON(w, http.StatusOK, res)
	}
}

func createDocumentCategory(serverCfg *api.ServerConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		type parameters struct {
			Name string `json:"name"`
		}
		params := parameters{}

		defer r.Body.Close()
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&params)
		if err != nil {
			api.RespondWithError(w, http.StatusBadRequest, "Payload error", err)
			return
		}

		doc_cat, err := serverCfg.DB.CreateDocumentCategory(r.Context(), params.Name)
		if err != nil {
			api.RespondWithError(w, http.StatusBadRequest, "Error adding to database", err)
			return
		}

		type response struct {
			ID        uuid.UUID `json:"id"`
			Name      string    `json:"name"`
			Active    bool      `json:"active"`
			CreatedAt time.Time `json:"created_at"`
			UpdatedAt time.Time `json:"updated_at"`
		}

		res := response{
			ID:        doc_cat.ID,
			Name:      doc_cat.Name,
			Active:    doc_cat.Active,
			CreatedAt: doc_cat.CreatedAt,
			UpdatedAt: doc_cat.UpdatedAt,
		}
		api.RespondWithJSON(w, http.StatusOK, res)
	}
}

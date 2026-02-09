package document_categories

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Mickdevv/savefuel-backend/api"
	"github.com/Mickdevv/savefuel-backend/internal/database"
	"github.com/google/uuid"
)

func GetDocumentCategories(serverCfg *api.ServerConfig) http.HandlerFunc {
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
func GetDocumentCategoryById(serverCfg *api.ServerConfig) http.HandlerFunc {
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

		api.RespondWithJSON(w, http.StatusOK, DocumentCategoryResponse{
			ID:        doc_cat.ID,
			Name:      doc_cat.Name,
			Active:    doc_cat.Active,
			CreatedAt: doc_cat.CreatedAt,
			UpdatedAt: doc_cat.UpdatedAt,
		})
	}

}

func DeleteDocumentCategory(serverCfg *api.ServerConfig) http.HandlerFunc {
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

func UpdateDocumentCategory(serverCfg *api.ServerConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		idStr := r.PathValue("id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			api.RespondWithError(w, http.StatusBadRequest, "Invalid ID format", err)
			return
		}

		params := UpdateDocumentCategoryPayload{}

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

		res := DocumentCategoryResponse{
			ID:        doc_cat.ID,
			Name:      doc_cat.Name,
			Active:    doc_cat.Active,
			CreatedAt: doc_cat.CreatedAt,
			UpdatedAt: doc_cat.UpdatedAt,
		}
		api.RespondWithJSON(w, http.StatusOK, res)
	}
}

func CreateDocumentCategory(serverCfg *api.ServerConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := CreateDocumentCategoryPayload{}

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

		res := DocumentCategoryResponse{
			ID:        doc_cat.ID,
			Name:      doc_cat.Name,
			Active:    doc_cat.Active,
			CreatedAt: doc_cat.CreatedAt,
			UpdatedAt: doc_cat.UpdatedAt,
		}
		api.RespondWithJSON(w, http.StatusOK, res)
	}
}

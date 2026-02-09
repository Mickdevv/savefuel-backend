package document_categories

import (
	"time"

	"github.com/google/uuid"
)

type CreateDocumentCategoryPayload struct {
	Name string `json:name`
}
type DocumentCategoryResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type UpdateDocumentCategoryPayload struct {
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

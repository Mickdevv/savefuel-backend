package documents

import (
	"time"

	"github.com/google/uuid"
)

type UploadDocumentPayload struct {
	Title       string    `json:"title"`
	Locale      string    `json:"locale"`
	Description string    `json:"description"`
	Priority    int32     `json:"priority"`
	CategoryID  uuid.UUID `json:"category_id"`
}
type DocumentPayload struct {
	Title       string    `json:"title"`
	Locale      string    `json:"locale"`
	Description string    `json:"description"`
	Priority    int32     `json:"priority"`
	Active      bool      `json:"active"`
	CategoryID  uuid.UUID `json:"category_id"`
}

type Document struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	Title       string    `json:"title"`
	Filename    string    `json:"filename"`
	Locale      string    `json:"locale"`
	Description string    `json:"description"`
	Priority    int32     `json:"priority"`
	Path        string    `json:"path"`
	Filetype    string    `json:"filetype"`
	Active      bool      `json:"active"`
	CategoryID  uuid.UUID `json:"category_id"`
}

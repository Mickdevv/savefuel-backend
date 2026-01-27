package documents

import (
	"net/http"

	"github.com/Mickdevv/savefuel-backend/api"
)

func RegisterRoutes(mux *http.ServeMux, serverCfg *api.ServerConfig) {
	mux.HandleFunc("GET /documents", api.WithCfg(serverCfg, getDocuments))
	mux.HandleFunc("POST /documents/upload", api.WithCfg(serverCfg, uploadDocument))
}

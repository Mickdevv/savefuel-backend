package documents

import (
	"github.com/Mickdevv/savefuel-backend/api"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, serverCfg *api.ServerConfig) {
	mux.HandleFunc("GET /documents", api.WithCfg(serverCfg, getDocuments))
	mux.HandleFunc("POST /documents/upload", api.WithCfg(serverCfg, uploadDocument))
}

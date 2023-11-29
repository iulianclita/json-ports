package app

import (
	"net/http"

	"github.com/iulianclita/json-ports/internal/presentation"
)

type route struct {
	path    string
	handler http.HandlerFunc
}

func (app *App) registerRoutes(mux *http.ServeMux) {
	routes := []route{
		{
			path:    "/ports",
			handler: app.portsHandler(),
		},
	}

	for _, route := range routes {
		mux.HandleFunc(route.path, route.handler)
	}
}

func (app *App) portsHandler() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		ports, err := presentation.ParseUpsertPortRequest(req)
		if err != nil {
			app.logger.Error("failed to parse upsert port request: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		responsePorts := presentation.PrepareUpsertPortResponse(ports)
		presentation.SendResponse(w, responsePorts, http.StatusCreated, app.logger)
	}
}

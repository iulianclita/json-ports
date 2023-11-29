package app

import (
	"fmt"
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
		switch req.Method {
		case http.MethodGet:
			ports, err := app.portService.GetPorts(req.Context())
			if err != nil {
				app.logger.Error("failed to get ports", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			responsePorts := presentation.PrepareGetPortsResponse(ports)
			presentation.SendResponse(w, responsePorts, http.StatusOK, app.logger)
		case http.MethodPost:
			ports, err := presentation.ParseUpsertPortRequest(req)
			if err != nil {
				app.logger.Error("failed to parse upsert port request", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			for _, port := range ports {
				if err := app.portService.UpsertPort(req.Context(), port); err != nil {
					app.logger.Error("failed to upsert port", fmt.Sprintf("port id = %s", port.ID), err)
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
					return
				}
			}

			presentation.SendResponse(w, nil, http.StatusCreated, app.logger)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	}
}

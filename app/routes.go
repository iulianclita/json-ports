package app

import "net/http"

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

	}
}

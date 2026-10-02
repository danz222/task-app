package web_transport_http

import (
	core_http_server "github.com/danz222/task-app/internal/core/transport/http/server"
)

type WebHttpHandler struct {
	webService WebService
}

type WebService interface {
	GetMainPage() ([]byte, error)
}

func NewWebHTTPHandler(
	webService WebService,
) *WebHttpHandler {
	return &WebHttpHandler{
		webService: webService,
	}
}

func (h *WebHttpHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Path:    "/",
			Handler: h.GetMainPage,
		},
	}
}

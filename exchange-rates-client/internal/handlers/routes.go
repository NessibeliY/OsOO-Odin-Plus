package handlers

import (
	"exchange-rates-client/nyeltay/internal/requestapi"
	"html/template"
	"net/http"
)

type Handler struct {
	cache      map[string]*template.Template
	apiRequest *requestapi.RequestAPI
}

func NewApplication(c map[string]*template.Template, r *requestapi.RequestAPI) *Handler {
	return &Handler{
		cache:      c,
		apiRequest: r,
	}
}

func (h *Handler) Routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./ui/static/"))))
	mux.Handle("/favicon.ico", http.NotFoundHandler())

	mux.HandleFunc("/", h.Home)
	mux.HandleFunc("/currencies", h.Url)

	return mux
}

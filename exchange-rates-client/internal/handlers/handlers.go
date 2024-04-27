package handlers

import (
	"exchange-rates-client/nyeltay/internal/models"
	"net/http"
	"net/url"
	"strconv"
)

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		h.notFound(w)
		return
	}

	if r.Method != http.MethodGet {
		h.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	err := h.apiRequest.Api()
	if err != nil {
		h.ServerError(w, err)
		return
	}

	filename := "index.html"

	h.render(w, filename, models.Currencies, http.StatusOK)
}

func (h *Handler) Url(w http.ResponseWriter, r *http.Request) {
	query, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		h.clientError(w, http.StatusBadRequest)
		return
	}

	if r.URL.Path != "/currencies" {
		h.notFound(w)
		return
	}

	if r.Method != http.MethodGet {
		h.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	id := query.Get("id")
	num, err := strconv.Atoi(id)
	if err != nil {
		h.clientError(w, http.StatusBadRequest)
		return
	}

	err = h.apiRequest.Api()
	if err != nil {
		h.ServerError(w, err)
		return
	}

	filename := "currency.html"

	h.render(w, filename, models.Currencies[num], http.StatusOK)
}

func (h *Handler) notFound(w http.ResponseWriter) {
	h.clientError(w, http.StatusNotFound)
}

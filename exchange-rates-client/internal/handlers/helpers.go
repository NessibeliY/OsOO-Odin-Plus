package handlers

import (
	"exchange-rates-client/nyeltay/internal/models"
	"fmt"
	"net/http"
)

func (h *Handler) clientError(w http.ResponseWriter, status int) {
	filename := "error.html"
	h.render(w, filename, &models.Error{
		ErrorText: http.StatusText(status),
		ErrorCode: status,
	}, status)
}

func (h *Handler) ServerError(w http.ResponseWriter, err error) {
	fmt.Println("server error:", err)

	filename := "error.html"

	h.render(w, filename, &models.Error{
		ErrorText: http.StatusText(http.StatusInternalServerError),
		ErrorCode: http.StatusInternalServerError,
	}, http.StatusInternalServerError)
}

func (h *Handler) render(w http.ResponseWriter, page string, data interface{}, status int) {
	ts, ok := h.cache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		h.ServerError(w, err)
		return
	}

	w.WriteHeader(status)
	err := ts.ExecuteTemplate(w, page, data)
	if err != nil {
		h.ServerError(w, err)
		return
	}
}

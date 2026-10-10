package account

import "net/http"

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Mount(mux *http.ServeMux) {
	mux.HandleFunc("GET /accounts/{id}", h.GetById)
}

func (h *Handler) GetById(w http.ResponseWriter, r *http.Request) {

}

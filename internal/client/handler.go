package client

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
	mux.HandleFunc("POST /client", h.NewClient)
	mux.HandleFunc("GET /client/{id}/accounts", h.GetAllAccounts)
	mux.HandleFunc("POST /client/{id}/accounts", h.NewAccount)
}

func (h *Handler) NewClient(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) GetAllAccounts(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) NewAccount(w http.ResponseWriter, r *http.Request) {

}

package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/todd-sudo/checker_server/pkg/logging"
)

type handler struct {
	log logging.Logger
}

type response struct {
	Ok  bool   `json:"ok"`
	Msg string `json:"msg"`
}

func NewHandler(log logging.Logger) *handler {
	return &handler{
		log: log,
	}
}

func (h *handler) InitRoutes(router *mux.Router) {
	router.HandleFunc("/heartbeat", h.checkServer).Methods("GET")
}

func (h *handler) checkServer(w http.ResponseWriter, r *http.Request) {
	res := response{
		Ok:  true,
		Msg: "I'm alive",
	}
	json.NewEncoder(w).Encode(res)
}

package server

import (
	"encoding/json"
	"net/http"
)

type GetRequestResponse struct {
	Data string `json:"data"`
}

func (s *Server) requestGet(w http.ResponseWriter, r *http.Request) {

	resp := GetRequestResponse{
		Data: s.DB.GetRandom(),
	}

	b, _ := json.Marshal(resp)

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

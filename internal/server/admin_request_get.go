package server

import (
	"encoding/json"
	"net/http"
	"taxiRequests/taxi_requests"
)

type GetAdminRequestResponse struct {
	ActiveRequests   [50]*taxi_requests.Request `json:"active_requests"`
	InactiveRequests []*taxi_requests.Request   `json:"inactive_requests"`
}

func (s *Server) adminRequestGet(w http.ResponseWriter, r *http.Request) {
	//
	resp := &GetAdminRequestResponse{
		ActiveRequests:   s.DB.ActiveReq.Requests,
		InactiveRequests: s.DB.InactiveReq.Requests,
	}

	b, _ := json.Marshal(resp)

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

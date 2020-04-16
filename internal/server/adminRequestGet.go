package server

import (
	"encoding/json"
	"net/http"
	"taxiRequests/taxiRequests"
)

type GetAdminRequestResponse struct {
	ActiveRequests   [50]*taxiRequests.Request `json:"active_requests"`
	InactiveRequests []*taxiRequests.Request   `json:"inactive_requests"`
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

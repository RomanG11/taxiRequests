package server

import (
	"encoding/json"
	"net/http"
)

type GetRequestResponse struct {
	Data string `json:"data"`
	A    int64
}

func (s *Server) requestGet(w http.ResponseWriter, r *http.Request) {

	var a int64
	for _, req := range s.DB.InactiveReq.Requests {
		a += req.Shown
	}

	for _, req := range s.DB.ActiveReq.Requests {
		a += req.Shown
	}

	resp := GetRequestResponse{
		Data: s.DB.GetRandom(),
		A:    a,
	}

	b, _ := json.Marshal(resp)

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

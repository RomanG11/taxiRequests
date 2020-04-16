package taxi_requests

import (
	"math/rand"
	strRand "taxiRequests/pkg/rand"
)

//
// ActiveReq contains active requests
//
type ActiveReq struct {
	Requests [50]*Request
}

//
// InitActReq initializing active requests
//
func InitActReq() *ActiveReq {
	actReq := ActiveReq{}

	for i := 0; i < 50; i++ {
		actReq.Requests[i] = &Request{
			Data:  strRand.String(2),
			Shown: 0,
		}
	}

	return &actReq
}

//
// GetRandom returns a random active request
//
func (ar *ActiveReq) GetRandom() string {
	r := rand.Int63n(50)

	ar.Requests[r].Shown++

	return ar.Requests[r].Data
}

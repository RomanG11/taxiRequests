package taxiRequests

import (
	"math/rand"
	strRand "taxiRequests/pkg/rand"
)

type ActiveReq struct {
	Requests [50]*Request
}

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

func (ar *ActiveReq) GetRandom() string {
	r := rand.Int63n(50)

	ar.Requests[r].Shown++

	return ar.Requests[r].Data
}

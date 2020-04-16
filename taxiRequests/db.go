package taxiRequests

import (
	"math/rand"
	strRand "taxiRequests/pkg/rand"
)

//
// DB contains all data requests
//
type DB struct {
	*ActiveReq
	*InactiveReq
}

//
// InitDB is a function to initialization taxi database
//
func InitDB() *DB {
	return &DB{
		ActiveReq:   InitActReq(),
		InactiveReq: &InactiveReq{},
	}
}

//
// Roll make disabled random active request and creating a new one
//
func (db *DB) Roll() {
	r := rand.Intn(50)

	if db.ActiveReq.Requests[r].Shown == 0 {
		return
	}

	db.InactiveReq.Requests = append(db.InactiveReq.Requests, db.ActiveReq.Requests[r])

	db.ActiveReq.Requests[r] = &Request{
		Data:  strRand.String(2),
		Shown: 0,
	}
}

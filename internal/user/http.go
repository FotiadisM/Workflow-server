package user

import (
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPHandler(e Endpoints, r *mux.Router, options ...httptransport.ServerOption) {

	r.Methods("GET").Path("").Handler(httptransport.NewServer(
		e.getUsersEndpoint,
		decodeGetUsersRequest,
		httptransport.EncodeJSONResponse,
		options...,
	))

	r.Methods("GET").Path("/{id}").Handler(httptransport.NewServer(
		e.getUserEndpoint,
		decodeGetUserRequest,
		httptransport.EncodeJSONResponse,
		options...,
	))

	r.Methods("GET").Path("/perpetrator/{id}").Handler(httptransport.NewServer(
		e.getPerpetatorEndpoint,
		decodeGetPerpetatorRequest,
		httptransport.EncodeJSONResponse,
		options...,
	))

	r.Methods("GET").Path("/connections/{id}").Handler(httptransport.NewServer(
		e.getConnectionsEndpoint,
		decodeGetConnectionsRequest,
		httptransport.EncodeJSONResponse,
		options...,
	))

	r.Methods("POST").Path("/connections").Handler(httptransport.NewServer(
		e.postConnectionEndpoint,
		decodePostConnectionRequest,
		httptransport.EncodeJSONResponse,
		options...,
	))

	r.Methods("PUT").Path("/connections").Handler(httptransport.NewServer(
		e.changeConnectionEndpoint,
		decodeChangeConnectionRequest,
		httptransport.EncodeJSONResponse,
		options...,
	))

	r.Methods("GET").Path("/connectionRequests/{id}").Handler(httptransport.NewServer(
		e.getConnectionRequestsEndpoint,
		decodeGetConnectionsRequest,
		httptransport.EncodeJSONResponse,
		options...,
	))

	r.Methods("PUT").Path("/connectionRequests").Handler(httptransport.NewServer(
		e.decideConnectionRequestEndpoint,
		decodeGetConnectionsRequests,
		httptransport.EncodeJSONResponse,
		options...,
	))
}

func decodeGetUsersRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	return getUsersRequest{}, nil
}

func decodeGetUserRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id := vars["id"]

	return getUserRequest{UserID: id}, nil
}

func decodeGetPerpetatorRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id := vars["id"]
	return getPerpetatorRequest{PerpID: id}, nil
}

func decodeGetConnectionsRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id := vars["id"]
	return getConnectionRequestsRequst{UserID: id}, nil
}

func decodePostConnectionRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	var res postConnectionRequest
	err = json.NewDecoder(r.Body).Decode(&res)

	return res, err
}

func decodeChangeConnectionRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	return
}

func decodeGetConnectionsRequests(ctx context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id := vars["id"]
	return getConnectionRequestsRequst{UserID: id}, nil
}

func decodeDecideConnectionRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	return
}

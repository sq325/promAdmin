package prom

import (
	"context"
	"encoding/json"
	"net/http"
)

func DecodeInstancesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

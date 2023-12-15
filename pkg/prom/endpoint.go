package prom

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type InstancesResponse struct {
	Instances []map[string]any `json:"instances"`
	Err       string           `json:"err"`
}

func MakePromSvcEndpoint(svc PromSvc) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		instances := svc.Instances()
		return InstancesResponse{Instances: instances}, nil
	}
}

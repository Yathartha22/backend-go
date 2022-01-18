package backend

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// type Endpoint func(ctx context.Context, request interface{}) (response interface{}, err error)

type Endpoints struct {
	GetInfoRequestEndpoint endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		GetInfoRequestEndpoint: MakeGetInfoRequestEndpoint(s),
	}
}

func MakeGetInfoRequestEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetInfoRequest)
		resp, err := s.GetInfo(ctx, req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

package backend

import "context"

type Service interface {
	GetInfo(ctx context.Context, request GetInfoRequest) (GetInfoResponse, error)
}

type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) GetInfo(ctx context.Context, req GetInfoRequest) (GetInfoResponse, error) {
	info := ""
	if req.Id == 1 {
		info = "Info for id 1"
	} else {
		info = "Info for some other id"
	}

	return GetInfoResponse{Info: info}, nil
}

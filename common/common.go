package common

import "github.com/arensusu/pionex/domain"

type CommonService struct {
	client domain.Client
}

func NewCommonService(c domain.Client) *CommonService {
	return &CommonService{client: c}
}

package main

import (
	advertise "advertiseproject/kitex_gen/advertiseproject/advertise"
	"advertiseproject/kitex_gen/base"
	"context"
	"errors"
	"log"
	"os"
	"time"
)

// AdServiceImpl implements the last service interface defined in the IDL.
type AdServiceImpl struct{}

// GetAd implements the AdServiceImpl interface.
func (s *AdServiceImpl) GetAd(ctx context.Context, req *advertise.GetAdReq) (resp *advertise.GetAdRes, err error) {
	if req == nil {
		return nil, errors.New("empty request")
	}

	// Simulate downstream timeout for degradation demo.
	if req.Name == "slow" || req.Id == 999 {
		time.Sleep(2 * time.Second)
	}

	// Simulate business failure for circuit-breaker demo.
	if req.Id < 0 {
		return nil, errors.New("simulated business error")
	}

	// Service-side degrade switch.
	if os.Getenv("AD_DEGRADE_MODE") == "1" || req.Name == "degrade" {
		resp = &advertise.GetAdRes{
			Ad: &advertise.Advertise{
				Id:          req.Id,
				Name:        "degraded-ad",
				Description: "fallback response from degrade mode",
				Stock:       0,
			},
			BaseRes: base.NewBaseRes(),
		}
		resp.BaseRes.Code = 0
		resp.BaseRes.Msg = "degraded response"
		return resp, nil
	}

	resp = &advertise.GetAdRes{}
	resp.Ad = &advertise.Advertise{
		Id:          req.Id,
		Name:        "fu",
		Description: "你访问对了",
		Stock:       999,
	}
	resp.BaseRes = base.NewBaseRes()
	resp.BaseRes.Code = 0
	resp.BaseRes.Msg = "ok"
	log.Println(resp)
	return
}

// AddAd implements the AdServiceImpl interface.
func (s *AdServiceImpl) AddAd(ctx context.Context, req *advertise.AddAdReq) (resp *advertise.AddAdRes, err error) {
	// TODO: Your code here...
	return
}

// DeleteAd implements the AdServiceImpl interface.
func (s *AdServiceImpl) DeleteAd(ctx context.Context, req *advertise.DeleteAdReq) (resp *advertise.DeleteAdRes, err error) {
	// TODO: Your code here...
	return
}

// UpdateAd implements the AdServiceImpl interface.
func (s *AdServiceImpl) UpdateAd(ctx context.Context, req *advertise.UpdateAdReq) (resp *advertise.UpdateAdRes, err error) {
	// TODO: Your code here...
	return
}

package main

import (
	advertise "advertiseproject/kitex_gen/advertiseproject/advertise"
	"advertiseproject/kitex_gen/base"
	"context"
	"log"
)

// AdServiceImpl implements the last service interface defined in the IDL.
type AdServiceImpl struct{}

// GetAd implements the AdServiceImpl interface.
func (s *AdServiceImpl) GetAd(ctx context.Context, req *advertise.GetAdReq) (resp *advertise.GetAdRes, err error) {
	// TODO: Your code here...
	resp = &advertise.GetAdRes{}
	resp.Ad = &advertise.Advertise{
		Id:          req.Id,
		Name:        "fuckyou",
		Description: "ffucku",
		Stock:       258,
	}
	resp.BaseRes = base.NewBaseRes()
	resp.BaseRes.Code = 0
	resp.BaseRes.Msg = "fffffffuuuuckuuuuu"
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

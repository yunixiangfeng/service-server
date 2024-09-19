package services

import (
	"context"
	// "git.internal.attains.cn/attains-cloud/go-api/attains/server/api"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Formal struct {
}

func (f Formal) ServerFormal(_ context.Context, request *api.ServerFormalRequest, response *api.ServerNewResponse) error {
	if data, err := Client().ServerFormal(request.Ip, request.Month); err != nil {
		return err
	} else {
		response.Id = data.Id
		response.Ip = data.Ip
		response.FreeDomain = data.FreeDomain
		response.Password = data.Password
		response.ExpireAt = timestamppb.New(data.ExpireAt)
	}
	return nil
}

func (f Formal) ServerGetFormalPrice(_ context.Context, request *api.ServerFormalRequest, response *api.ServerGetFormalPriceResponse) error {
	if price, err := Client().ServerGetFormalPrice(request.Ip, request.Month); err != nil {
		return err
	} else {
		response.Price = price
	}
	return nil
}

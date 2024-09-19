package services

import (
	"context"
	// "git.internal.attains.cn/attains-cloud/go-api/attains/server/api"
)

type HostSnapSetDiy struct {
}

func (h HostSnapSetDiy) ServerHostSnapSetDiy(_ context.Context, request *api.ServerHostSnapSetDiyRequest, response *api.ServerStringListResponse) error {
	if data, err := Client().ServerHostSnapSetDiy(request.Ip, request.Hours, request.Weeks, request.KeepDay); err != nil {
		return err
	} else {
		response.Result = data
	}
	return nil
}

func (h HostSnapSetDiy) ServerHostSnapGetCancelDiy(_ context.Context, request *api.ServerRequest, response *api.ServerHostSnapGetCancelDiyResponse) error {
	if data, err := Client().ServerHostSnapGetcanceldiy(request.Ip); err != nil {
		return err
	} else {
		response.Size = data.Size
		response.Price = int64(data.Price * 100)
		response.NoPaySize = data.Nopaysize
		response.TotalPrice = int64(data.Totalprice * 100)
	}
	return nil
}

func (h HostSnapSetDiy) ServerHostSnapCancelDiy(_ context.Context, request *api.ServerHostSnapSetDiyRequest, response *api.ServerStringListResponse) error {
	if data, err := Client().ServerHostSnapCanceldiy(request.Ip); err != nil {
		return err
	} else {
		response.Result = data
	}
	return nil
}

func (h HostSnapSetDiy) ServerHostSnapDiySnapCharge(_ context.Context, request *api.ServerHostSnapSetDiyRequest, response *api.ServerStringListResponse) error {
	if data, err := Client().ServerHostSnapDiysnapcharge(request.Ip); err != nil {
		return err
	} else {
		response.Result = data
	}
	return nil
}

package services

import (
	"context"
	// commonApi "git.internal.attains.cn/attains-cloud/go-api/attains/common/api"
	// serverApi "git.internal.attains.cn/attains-cloud/go-api/attains/server/api"
)

type Status struct {
}

func (s Status) ServerSetStatus(_ context.Context, request *serverApi.ServerSetStatusRequest, _ *commonApi.EmptyResponse) error {
	if err := Client().ServerSetStatus(request.Ip, request.Status); err != nil {
		return err
	}
	return nil
}

func (s Status) ServerGetStatus(_ context.Context, request *serverApi.ServerRequest, response *serverApi.ServerGetStatusResponse) error {
	if data, err := Client().ServerGetStatus(request.Ip); err != nil {
		return err
	} else {
		response.Status = data.Status
		response.Info = data.Info
	}
	return nil
}

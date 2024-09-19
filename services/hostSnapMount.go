package services

import (
	"context"
	// "git.internal.attains.cn/attains-cloud/go-api/attains/server/api"
)

type HostSnapMount struct {
}

func (h HostSnapMount) ServerHostSnapMountList(_ context.Context, request *api.ServerRequest, response *api.ServerHostSnapMountListResponse) error {
	if data, err := Client().ServerHostSnapMountList(request.Ip); err != nil {
		return err
	} else {
		response.SnapMountList.Os = data.Snapmountlist.Os
		response.SnapMountList.Data = data.Snapmountlist.Data
	}
	return nil
}

func (h HostSnapMount) ServerHostSnapMount(_ context.Context, request *api.ServerHostSnapMountRequest, response *api.ServerStringListResponse) error {
	if data, err := Client().ServerHostSnapMount(request.Ip, request.Os, request.Data); err != nil {
		return err
	} else {
		response.Result = data
	}
	return nil
}

func (h HostSnapMount) ServerHostSnapUnmount(_ context.Context, request *api.ServerHostSnapUnmountRequest, response *api.ServerStringListResponse) error {
	if data, err := Client().ServerHostSnapUnmount(request.Ip, request.Info); err != nil {
		return err
	} else {
		response.Result = data
	}
	return nil
}

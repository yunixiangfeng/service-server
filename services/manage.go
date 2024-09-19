package services

import (
	"context"
	// commonApi "git.internal.attains.cn/attains-cloud/go-api/attains/common/api"
	// serverApi "git.internal.attains.cn/attains-cloud/go-api/attains/server/api"
)

type Manage struct {
}

func (m Manage) ServerChangePassword(_ context.Context, request *serverApi.ServerChangePasswordRequest, _ *commonApi.EmptyResponse) error {
	if err := Client().ServerChangePassword(request.Ip, request.NewPwd, request.PwdType); err != nil {
		return err
	}
	return nil
}

func (m Manage) ServerDenyPanelLogin(_ context.Context, request *serverApi.ServerRequest, _ *commonApi.EmptyResponse) error {
	if err := Client().ServerDenyPanelLogin(request.Ip); err != nil {
		return err
	}
	return nil
}

func (m Manage) ServerGetPassword(_ context.Context, request *serverApi.ServerRequest, response *serverApi.ServerGetPasswordResponse) error {
	if password, err := Client().ServerGetPassword(request.Ip); err != nil {
		return err
	} else {
		response.Password = password
	}
	return nil
}

func (m Manage) ServerReinstall(_ context.Context, request *serverApi.ServerReinstallRequest, _ *commonApi.EmptyResponse) error {
	if err := Client().ServerReinstall(request.Ip, request.NewOs, request.KeepData); err != nil {
		return err
	}
	return nil
}

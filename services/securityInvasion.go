package services

import (
	"context"
	// "git.internal.attains.cn/attains-cloud/go-api/attains/server/api"
)

type SecurityInvasion struct {
}

func (s SecurityInvasion) ServerSecurityInvasionStatus(_ context.Context, request *api.ServerRequest, response *api.ServerSecurityInvasionStatusResponse) error {
	if status, err := Client().ServerSecurityInvasionStatus(request.Ip); err != nil {
		return err
	} else {
		response.Status = status
	}
	return nil
}

func (s SecurityInvasion) ServerSecurityInvasionLogs(_ context.Context, request *api.ServerListRequest, response *api.ServerSecurityInvasionLogsResponse) error {
	if data, err := Client().ServerSecurityInvasionLogs(request.Ip, request.PageNo, request.Limit); err != nil {
		return err
	} else {
		response.Count = data.Count
		response.Total = data.Total
		response.Count7 = data.Count7
		response.Count30 = data.Count30
		response.Items = data.Items
	}
	return nil
}

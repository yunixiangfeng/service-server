package services

import (
	"context"
	// commonApi "git.internal.attains.cn/attains-cloud/go-api/attains/common/api"
	// serverApi "git.internal.attains.cn/attains-cloud/go-api/attains/server/api"
)

type Firewall struct {
}

func (f Firewall) ServerFirewallList(_ context.Context, request *serverApi.ServerListRequest, response *serverApi.ServerFirewallListResponse) error {
	if data, err := Client().ServerFirewallList(request.Ip, request.PageNo, request.Limit); err != nil {
		return err
	} else {
		response.Total = data.Total
		response.Limit = data.Limit
		response.PageNo = data.Pageno
		response.PageCount = data.Pagecount
		for _, v := range data.Items {
			response.List = append(response.List, &serverApi.FirewallList{
				SysId:     v.Sysid,
				GroupName: v.GroupName,
				GroupMemo: v.GroupMemo,
				State:     v.State,
				CreateAt:  v.CreateAt,
				IsSystem:  v.Issystem,
				Count:     v.Count,
			})
		}
	}
	return nil
}

func (f Firewall) ServerFirewallAdd(_ context.Context, request *serverApi.ServerFirewallAddRequest, _ *commonApi.EmptyResponse) error {
	if err := Client().ServerFirewallAdd(request.Name, request.Memo); err != nil {
		return err
	}
	return nil
}

func (f Firewall) ServerFirewallDel(_ context.Context, request *serverApi.ServerIdlRequest, response *serverApi.ServerStringListResponse) error {
	if data, err := Client().ServerFirewallDel(request.Id); err != nil {
		return err
	} else {
		response.Result = data
	}
	return nil
}

func (f Firewall) ServerFirewallSet(_ context.Context, request *serverApi.ServerFirewallSetRequest, response *serverApi.ServerStringListResponse) error {
	if data, err := Client().ServerFirewallGroupSet(request.GroupId, request.Sip, request.ActCmd); err != nil {
		return err
	} else {
		response.Result = data
	}
	return nil
}

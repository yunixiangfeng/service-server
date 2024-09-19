package services

import (
	"context"
	// serverRequest "git.internal.attains.cn/attains-cloud/cloud-provider-sdk-west/request"
	// "git.internal.attains.cn/attains-cloud/go-api/attains/server/api"
)

type FirewallRule struct {
}

func (f FirewallRule) ServerFirewallGroupRules(_ context.Context, request *api.ServerIdlRequest, response *api.ServerFirewallGroupRulesResponse) error {
	if data, err := Client().ServerFirewallGroupRules(request.Id); err != nil {
		return err
	} else {
		for _, v := range data {
			response.List = append(response.List, &api.FirewallGroupRules{
				SysId:      v.Sysid,
				GroupId:    v.GroupId,
				Proto:      v.Proto,
				SrcIp:      v.Srcip,
				SrcPort:    v.Srcport,
				DstPort:    v.Dstport,
				Priority:   v.Priority,
				Action:     v.Action,
				Remark:     v.Remark,
				SrcPortMax: v.SrcportMax,
				DstPortMax: v.DstportMax,
			})
		}
	}
	return nil
}

func (f FirewallRule) ServerFirewallGroupRuleAdd(_ context.Context, request *api.ServerFirewallGroupRuleAddRequest, response *api.ServerStringListResponse) error {
	req := &serverRequest.ServerFirewallGroupRuleAdd{
		Groupid:    request.GroupId,
		Proto:      request.Proto,
		SrcipType:  request.SrcIpType,
		Srcip:      request.SrcIp,
		Porttype:   request.PortType,
		Srcport:    request.SrcPort,
		SrcportMax: request.SrcPortMax,
		Dstport:    request.DstPort,
		DstportMax: request.DstPortMax,
		Priority:   request.Priority,
		Action:     request.Action,
	}
	if data, err := Client().ServerFirewallGroupRuleAdd(req); err != nil {
		return err
	} else {
		response.Result = data
	}
	return nil
}

func (f FirewallRule) ServerFirewallGroupRuleDel(_ context.Context, request *api.ServerFirewallGroupRuleDelRequest, response *api.ServerStringListResponse) error {
	if data, err := Client().ServerFirewallGroupRuleDel(request.GroupId, request.RuleId); err != nil {
		return err
	} else {
		response.Result = data
	}
	return nil
}

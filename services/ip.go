package services

import (
	"context"
	// commonApi "git.internal.attains.cn/attains-cloud/go-api/attains/common/api"
	// serverApi "git.internal.attains.cn/attains-cloud/go-api/attains/server/api"
)

type Ip struct {
}

func (i Ip) ServerGetIpById(_ context.Context, request *serverApi.ServerIdlRequest, response *serverApi.ServerGetIpByIdResponse) error {
	if ip, err := Client().ServerGetIpById(request.Id); err != nil {
		return err
	} else {
		response.Ip = ip
	}
	return nil
}

func (i Ip) ServerCanBuyOtherIp(_ context.Context, request *serverApi.ServerRequest, response *serverApi.ServerCanBuyOtherIpResponse) error {
	if data, err := Client().ServerCanBuyOtherIp(request.Ip); err != nil {
		return err
	} else {
		response.Count = data.Count
		response.Max = data.Max
		response.Used = data.Used
		for _, v := range data.Iptype {
			response.IpType = append(response.IpType, &serverApi.ServerCanBuyOtherIpType{
				Value: v.Value,
				Name:  v.Name,
				Count: v.Count,
			})
		}
	}
	return nil
}

func (i Ip) ServerAddIp(_ context.Context, request *serverApi.ServerIpRequest, _ *commonApi.EmptyResponse) error {
	if err := Client().ServerAddIp(request.Ip, request.IpType, request.BuyCount); err != nil {
		return err
	}
	return nil
}

func (i Ip) ServerGetIpv6(_ context.Context, request *serverApi.ServerRequest, response *serverApi.ServerGetIpv6Response) error {
	if data, err := Client().ServerGetIpv6(request.Ip); err != nil {
		return err
	} else {
		response.Ipv6Gateway = data.Ipv6gateway
		response.IsAddIpv6 = data.Isaddipv6
		response.IsUseIpv6 = data.Isuseipv6
		response.Ipv6Dns = data.Ipv6dns
		response.AllIpv6 = data.AllIpv6
	}
	return nil
}

func (i Ip) ServerGetOtherIpPrice(_ context.Context, request *serverApi.ServerIpRequest, response *serverApi.ServerGetOtherIpPriceResponse) error {
	if price, err := Client().ServerGetOtherIpPrice(request.Ip, request.IpType, request.BuyCount); err != nil {
		return err
	} else {
		response.Price = price
	}
	return nil
}

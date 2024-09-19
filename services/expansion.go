package services

import (
	"context"
	// serverRequest "git.internal.attains.cn/attains-cloud/cloud-provider-sdk-west/request"
	// commonApi "git.internal.attains.cn/attains-cloud/go-api/attains/common/api"
	// serverApi "git.internal.attains.cn/attains-cloud/go-api/attains/server/api"
)

type Expansion struct {
}

func (e Expansion) ServerExpansionList(_ context.Context, request *serverApi.ServerRequest, response *serverApi.ServerExpansionListResponse) error {
	if data, err := Client().ServerExpansionList(request.Ip); err != nil {
		return err
	} else {
		response.Total = data.Total
		response.Limit = data.Limit
		response.PageNo = data.Pageno
		response.PageCount = data.Pagecount
		for _, v := range data.Items {
			response.List = append(response.List, &serverApi.ExpansionList{
				Id:       v.Id,
				StarDay:  v.Starday,
				EndDay:   v.Endday,
				Status:   v.Status,
				Val:      v.Val,
				AddTime:  v.Addtime,
				TypeName: v.Typename,
			})
		}
	}
	return nil
}

func (e Expansion) ServerExpansionConfig(_ context.Context, request *serverApi.ServerRequest, response *serverApi.ServerExpansionConfigResponse) error {
	if data, err := Client().ServerExpansionConfig(request.Ip); err != nil {
		return err
	} else {
		for _, v := range data.Cpu {
			response.Cpu = append(response.Cpu, &serverApi.ExpansionConfigCpu{
				Size:  v.Size,
				Price: int64(v.Price * 100),
			})
		}
		for _, v := range data.Ram {
			response.Ram = append(response.Ram, &serverApi.ExpansionConfigRam{
				Size:  v.Size,
				Price: int64(v.Price * 100),
			})
		}
		response.Flux = &serverApi.ExpansionConfigFlux{Max: data.Flux.Max}
		response.Info = &serverApi.ExpansionConfigInfo{
			EndDate:  data.Info.Enddate,
			StarDate: data.Info.Stardate,
			Flux:     data.Info.Flux,
		}
	}
	return nil
}

func (e Expansion) ServerExpansionFluxPrice(_ context.Context, request *serverApi.ServerExpansionFluxRequest, response *serverApi.ServerExpansionFluxPriceResponse) error {
	req := &serverRequest.ServerExpansionFluxPrice{
		Ip:      request.Ip,
		Starday: request.StarDay,
		Endday:  request.EndDay,
		Flux:    request.Flux,
	}
	if data, err := Client().ServerExpansionFluxPrice(req); err != nil {
		return err
	} else {
		response.Total = data.Total
		response.Day = data.Day
		response.DayPrice = data.Dayprice
		response.Discount = data.Discount
	}
	return nil
}

func (e Expansion) ServerExpansionFlux(_ context.Context, request *serverApi.ServerExpansionFluxRequest, response *serverApi.ServerStringListResponse) error {
	req := &serverRequest.ServerExpansionFluxPrice{
		Ip:      request.Ip,
		Starday: request.StarDay,
		Endday:  request.EndDay,
		Flux:    request.Flux,
	}
	if data, err := Client().ServerExpansionFlux(req); err != nil {
		return err
	} else {
		response.Result = data
	}
	return nil
}

func (e Expansion) ServerExpansionUpdate(_ context.Context, request *serverApi.ServerExpansionUpdateRequest, _ *commonApi.EmptyResponse) error {
	req := &serverRequest.ServerExpansionUpdate{
		Ip:     request.Ip,
		Cpu:    request.Cpu,
		Ram:    request.Ram,
		Endday: request.EndDay,
	}
	if err := Client().ServerExpansionUpdate(req); err != nil {
		return err
	}
	return nil
}

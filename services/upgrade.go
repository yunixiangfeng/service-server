package services

import (
	"context"
	// serverRequest "git.internal.attains.cn/attains-cloud/cloud-provider-sdk-west/request"
	// commonApi "git.internal.attains.cn/attains-cloud/go-api/attains/common/api"
	// serverApi "git.internal.attains.cn/attains-cloud/go-api/attains/server/api"
)

type Upgrade struct {
}

func (u Upgrade) ServerGetUpgradeInfo(_ context.Context, request *serverApi.ServerRequest, response *serverApi.ServerGetUpgradeInfoResponse) error {
	if data, err := Client().ServerGetUpgradeInfo(request.Ip); err != nil {
		return err
	} else {
		response.Ip = data.Ip
		response.FreeDomain = data.Freedomain
		response.IsUpgrade = data.Isupgrade
		response.Cpu = data.Cpu
		response.Ram = data.Ram
		response.Flux = data.Flux
		response.OsData = data.Osdata
		response.Data = data.Data
		response.MinData = data.Mindata
		response.Room = data.Room
		response.Ddos = data.Ddos
		response.DiskType = data.Disktype
		response.MaxFlux = data.Maxflux
		response.MaxOsData = data.Maxosdata
		response.MaxData = data.Maxdata
		response.EbsName = data.Ebsname
	}
	return nil
}

func (u Upgrade) ServerGetUpgradePrice(_ context.Context, request *serverApi.ServerGetUpgradeRequest, response *serverApi.ServerGetUpgradePriceResponse) error {
	req := &serverRequest.ServerGetUpgradePrice{
		Ip:       request.Ip,
		Cpu:      request.Cpu,
		Ram:      request.Ram,
		Flux:     request.Flux,
		Osdata:   request.OsData,
		Data:     request.Data,
		Room:     request.Room,
		Disktype: request.DiskType,
		Blday:    request.BlDay,
		Ddos:     request.Ddos,
	}
	if data, err := Client().ServerGetUpgradePrice(req); err != nil {
		return err
	} else {
		response.DayPrice = int64(data.Dayprice * 100)
		response.LeftDay = data.Leftday
		response.Service = int64(data.Service * 100)
		response.UpPrice = int64(data.Upprice * 100)
		response.BlPrice = int64(data.Blprice * 100)
		response.FullPrice = int64(data.Fullprice * 100)
		response.NewMonthPrice = int64(data.Newmonthprice * 100)
		response.OldMonthPrice = int64(data.Oldmonthprice * 100)
		response.BillType = data.Billtype
		response.CalPrice = int64(data.Calprice * 100)
		response.ChgRoom = data.Chgroom
	}
	return nil
}

func (u Upgrade) ServerUpgrade(_ context.Context, request *serverApi.ServerGetUpgradeRequest, _ *commonApi.EmptyResponse) error {
	req := &serverRequest.ServerGetUpgradePrice{
		Ip:       request.Ip,
		Cpu:      request.Cpu,
		Ram:      request.Ram,
		Flux:     request.Flux,
		Osdata:   request.OsData,
		Data:     request.Data,
		Room:     request.Room,
		Disktype: request.DiskType,
		Blday:    request.BlDay,
		Ddos:     request.Ddos,
	}
	if err := Client().ServerUpgrade(req); err != nil {
		return err
	}
	return nil
}

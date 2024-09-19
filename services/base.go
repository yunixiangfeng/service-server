package services

import (
	"context"
	// clientRequest "git.internal.attains.cn/attains-cloud/cloud-provider-sdk-west/request"
	// "git.internal.attains.cn/attains-cloud/go-api/attains/server/api"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Base struct {
}

func (b Base) ServerCalPrice(_ context.Context, request *api.ServerCalPriceRequest, response *api.ServerCalPriceResponse) error {
	req := new(clientRequest.ServerCalPrice)
	req.CpuType = request.CpuType
	req.CpuNum = request.CpuNum
	req.Ram = request.Ram
	req.OsDisk = request.OsDisk
	req.DataDisk = request.DataDisk
	req.Flux = request.Flux
	req.Room = request.Room
	req.Ddos = request.Ddos
	req.DiskType = request.DiskType
	req.SnapAdv = request.SnapAdv
	req.CC = request.CC
	req.Month = request.Month
	req.Trial = request.Trial
	if data, err := Client().ServerCalPrice(req); err != nil {
		return err
	} else {
		response.Price = int64(data.Price * 100)
		response.FullPrice = int64(data.FullPrice * 100)
		response.ZtCount = data.ZtCount
		response.WzCount = data.WzCount
	}
	return nil
}

func (b Base) ServerNew(_ context.Context, request *api.ServerNewRequest, response *api.ServerNewResponse) error {
	req := new(clientRequest.ServerNew)
	req.CpuType = request.CpuType
	req.CpuNum = request.CpuNum
	req.Ram = request.Ram
	req.OsDisk = request.OsDisk
	req.DataDisk = request.DataDisk
	req.Flux = request.Flux
	req.Room = request.Room
	req.Ddos = request.Ddos
	req.DiskType = request.DiskType
	req.SnapAdv = request.SnapAdv
	req.CC = request.CC
	req.Month = request.Month
	req.Trial = request.Trial
	req.PriceMoney = float64(request.PriceMoney / 100)
	req.Os = request.Os
	req.Setebsid = request.SetEbsId
	if data, err := Client().ServerNew(req); err != nil {
		return err
	} else {
		response.Id = data.Id
		response.Ip = data.Ip
		response.FreeDomain = data.FreeDomain
		response.Password = data.Password
		response.ExpireAt = timestamppb.New(data.ExpireAt)
	}
	return nil
}

func (b Base) ServerGetList(_ context.Context, request *api.ServerGetListRequest, response *api.ServerGetListResponse) error {
	if data, err := Client().ServerGetList(request.PageNo, request.Limit, request.ServerRoom); err != nil {
		return err
	} else {
		response.PageNo = data.Pageno
		response.Limit = data.Limit
		response.PageCount = data.Pagecount
		response.Total = data.Total
		for _, v := range data.Items {
			response.List = append(response.List, &api.MyServerList{
				Id:          v.Id,
				Os:          v.Os,
				Cpu:         v.Cpu,
				OsHardDisk:  v.Osharddisk,
				HardDisk:    v.Harddisk,
				Memory:      v.Memory,
				MainBoard:   v.Mainboard,
				AllocateIp:  v.Allocateip,
				StartTime:   v.Starttime,
				ExpDate:     v.Expdate,
				ServerRoom:  v.Serverroom,
				AddedServer: v.Addedserver,
				ProId:       v.P_proid,
				Flux:        v.Flux,
				BuyTest:     v.Buytest,
				GroupId:     v.GroupId,
				DenyLogin:   v.Denylogin,
				FireGroup:   v.Firegroup,
				SnapAdv:     v.Snapadv,
				SnapAdvTime: v.Snapadvtime,
				Ddos:        v.Ddos,
				ProdType:    v.Prodtype,
				Cc:          v.Cc,
				LswId:       v.Lswid,
				MeterType:   v.Metertype,
				BillType:    v.Billtype,
				HkIp:        v.Hkip,
				EbsId:       v.Ebsid,
			})
		}
	}
	return nil
}

func (b Base) ServerDetail(_ context.Context, request *api.ServerRequest, response *api.ServerDetailResponse) error {
	if data, err := Client().ServerDetail(request.Ip); err != nil {
		return err
	} else {
		response.Id = data.Id
		response.AllocateIp = data.Allocateip
		response.ProdType = data.Prodtype
		response.Os = data.Os
		response.Cpu = data.Cpu
		response.OsHardDisk = data.Osharddisk
		response.HardDisk = data.Harddisk
		response.Memory = data.Memory
		response.Ddos = data.Ddos
		response.Cc = data.Cc
		response.StartTime = data.Starttime
		response.ServerRoom = data.Serverroom
		response.ExpTime = data.Exptime
		response.RamDomPass = data.Ramdompass
		response.AddedServer = data.Addedserver
		response.State = data.Sstate
		response.StateInfo = data.Sstateinfo
		response.ProId = data.P_proid
		response.FreeDomain = data.Freedomain
		response.Flux = data.Flux
		response.BuyTest = data.Buytest
		response.DiskType = data.Disktype
		response.DenyLogin = data.Denylogin
		response.FireGroup = data.Firegroup
		response.SnapAdv = data.Snapadv
		response.SnapPeriod = data.Snapperiod
		response.SnapAdvTime = data.Snapadvtime
		response.LswId = data.Lswid
		response.BillType = data.Billtype
		response.HkIp = data.Hkip
		response.NetGate = data.Netgate
		response.Netmask = data.Netmask
		response.Netmask = data.Netvlan
		response.SelfIp = data.V_selfip
		response.EbsId = data.Ebsid
		response.EbsName = data.Ebsname
		response.Port = data.Port
		response.Login = data.Login
		response.OtherIp = data.Otherip
		response.Extra = &api.ServerDetailExtra{
			AddFlux: data.Extra.Addflux,
			AddCpu:  data.Extra.Addcpu,
			AddRam:  data.Extra.Addram,
		}
		for k, v := range data.OsExtend {
			response.OsExtend[k] = &api.ServerDetailOsExtend{
				Name:  v.Name,
				Value: v.Value,
			}
		}
	}
	return nil
}

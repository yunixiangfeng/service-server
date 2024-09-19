package services

import (
	"context"
	"encoding/json"
	// serverRequest "git.internal.attains.cn/attains-cloud/cloud-provider-sdk-west/request"
	// "git.internal.attains.cn/attains-cloud/go-api/attains/server/api"
)

type HostSnap struct {
}

func (h HostSnap) ServerGetHostSnapList(_ context.Context, request *api.ServerRequest, response *api.ServerGetHostSnapListResponse) error {
	if data, err := Client().ServerGetHostSnapList(request.Ip); err != nil {
		return err
	} else {
		response.OS = data.Os
		response.Data = data.Data
		response.SnapPeriod = data.Snapperiod
		response.SnapAdv = data.Snapadv
		response.SnapAdvTime = data.Snapadvtime
		response.LeftMonth = data.Leftmonth
		response.ExpireDate = data.Expiredate
		strValue, _ := json.Marshal(data.Showday)
		response.ShowDay = string(strValue)
		response.DiySize = &api.HostSnapDiySize{
			Size:      data.Diysize.Size,
			NoPaySize: data.Diysize.Nopaysize,
		}
		response.SnapRules = &api.HostSnapRules{
			KeepDay:   data.Snaprules.Keepday,
			Hours:     data.Snaprules.Hours,
			Weeks:     data.Snaprules.Weeks,
			NoPaySize: data.Snaprules.Nopaysize,
		}
		response.SnapAdvPrice = int64(data.Snapadvprice * 100)
		response.SnapPolicy = data.Snappolicy
		response.IsCharge = data.Ischarge
		response.SnapRestorePrice = int64(data.SnapRestorePrice * 100)
		response.ProtectSnapPrice = int64(data.Protectsnapprice * 100)
		response.ProtectSnap = &api.HostSnapProtectSnap{
			SnapId: data.Protectsnap.Snapid,
			Time:   data.Protectsnap.Time,
		}
	}
	return nil
}

func (h HostSnap) ServerHostSnapBuySnapAdv(_ context.Context, request *api.ServerHostSnapBuySnapAdvRequest, response *api.ServerHostSnapBuySnapAdvResponse) error {
	if data, err := Client().ServerHostSnapBuySnapAdv(request.Ip, request.SnapAdvVal); err != nil {
		return err
	} else {
		response.SnapPeriod = data.Snapperiod
		response.SnapAdv = data.Snapadv
		response.SnapAdvTime = data.Snapadvtime
		response.LeftMonth = data.Leftmonth
		response.ExpireDate = data.Expiredate
		strValue, _ := json.Marshal(data.Showday)
		response.ShowDay = string(strValue)
		response.SnapAdvPrice = int64(data.Snapadvprice * 100)
		response.SnapPolicy = data.Snappolicy
		response.IsCharge = data.Ischarge
		response.SnapRestorePrice = int64(data.SnapRestorePrice * 100)
		response.SnapRules = &api.HostSnapAdvRules{
			KeepDay: data.Snaprules.Keepday,
			Hours:   data.Snaprules.Hours,
			Weeks:   data.Snaprules.Weeks,
		}
		for _, v := range data.Os {
			response.OS = append(response.OS, &api.HostSnapBuySnapAdv{
				Index: v.Index,
				Size:  v.Size,
				Date:  v.Date,
			})
		}
		for _, v := range data.Data {
			response.Data = append(response.Data, &api.HostSnapBuySnapAdv{
				Index: v.Index,
				Size:  v.Size,
				Date:  v.Date,
			})
		}
	}
	return nil
}

func (h HostSnap) ServerHostSnapRestore(_ context.Context, request *api.ServerHostSnapRestoreRequest, response *api.ServerStringListResponse) error {
	req := &serverRequest.ServerHostSnapRestore{
		Ip:       request.Ip,
		Os:       request.Os,
		Ostime:   request.OsTime,
		Data:     request.Data,
		Datatime: request.DataTime,
	}
	if data, err := Client().ServerHostSnapRestore(req); err != nil {
		return err
	} else {
		response.Result = data
	}
	return nil
}

func (h HostSnap) ServerHostSnapPeriod(_ context.Context, request *api.ServerHostSnapPeriodRequest, response *api.ServerStringListResponse) error {
	if data, err := Client().ServerHostSnapPeriod(request.Ip, request.SnapPeriod); err != nil {
		return err
	} else {
		response.Result = data
	}
	return nil
}

func (h HostSnap) ServerHostSnapManual(_ context.Context, request *api.ServerHostSnapManualRequest, response *api.ServerStringListResponse) error {
	if data, err := Client().ServerHostSnapBkup(request.Ip, request.Info); err != nil {
		return err
	} else {
		response.Result = data
	}
	return nil
}

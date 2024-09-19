package services

import (
	"context"
	// commonApi "git.internal.attains.cn/attains-cloud/go-api/attains/common/api"
	// serverApi "git.internal.attains.cn/attains-cloud/go-api/attains/server/api"
)

type Renew struct {
}

func (r Renew) ServerRenew(_ context.Context, request *serverApi.ServerRenewRequest, _ *commonApi.EmptyResponse) error {
	if err := Client().ServerRenew(request.Ip, request.Month); err != nil {
		return err
	}
	return nil
}

func (r Renew) ServerCalRenewPrice(_ context.Context, request *serverApi.ServerCalRenewPriceRequest, response *serverApi.ServerCalRenewPriceResponse) error {
	if data, err := Client().ServerCalRenewPrice(request.Ip, request.Month, request.SnapAdv); err != nil {
		return err
	} else {
		response.Price = int64(data.Price * 100)
		response.FluxPrice = int64(data.Fluxprice * 100)
		response.Info = data.Info
		response.YearCaTxt = data.Yearcatxt
		response.PreDay = data.Preday
		response.AddIpTxt = data.Addiptxt
		response.FullPrice = int64(data.Fullprice * 100)
		response.IsSnap = data.Issnap
		response.SnapPrice = int64(data.Snapprice * 100)
		response.SnapBuyPayMonth = data.Snapbuypaymonth
		response.SnapNoMoneyMonth = data.Snapnomoneymonth
		response.CcPrice = int64(data.Ccpirce * 100)
		response.RenewMonth = data.Renewmonth
		response.AddIpPrice = int64(data.Addipprice * 100)
		response.UseMoney = int64(data.Usemoney * 100)
	}
	return nil
}

package services

import (
	"context"
	// commonApi "git.internal.attains.cn/attains-cloud/go-api/attains/common/api"
	// serverApi "git.internal.attains.cn/attains-cloud/go-api/attains/server/api"
)

type White struct {
}

func (w White) ServerWhiteIpList(_ context.Context, request *serverApi.ServerRequest, response *serverApi.ServerWhiteIpListResponse) error {
	if data, err := Client().ServerWhiteIpList(request.Ip); err != nil {
		return err
	} else {
		response.BuyPrice = int64(data.Buyprice * 100)
		response.BuyZtPrice = int64(data.Buyztprice * 100)
		response.MaxBeiAn = data.Maxbeian
		response.BuyBeiAn = data.Buybeian
		response.MaxDomain = data.Maxdomain
		for _, v := range data.List {
			response.List = append(response.List, &serverApi.WhiteIpList{
				Ip:        v.Ip,
				BaseCount: v.Basecount,
				Count:     v.Count,
				AddCount:  v.Addcount,
				Used:      v.Used,
			})
		}
	}
	return nil
}

func (w White) ServerGetWhiteList(_ context.Context, request *serverApi.ServerListRequest, response *serverApi.ServerGetWhiteListResponse) error {
	if data, err := Client().ServerGetWhiteList(request.Ip, request.PageNo, request.Limit); err != nil {
		return err
	} else {
		response.Total = data.Total
		response.Limit = data.Limit
		response.PageNo = data.Pageno
		response.PageCount = data.Pagecount
		for _, v := range data.Items {
			response.List = append(response.List, &serverApi.GetWhiteList{
				RowNumber: v.Rownumber,
				Id:        v.Id,
				Ip:        v.Ip,
				StrDomain: v.Strdomain,
				IcpNo:     v.Icpno,
				WebTitle:  v.Webtitle,
				IsJieRu:   v.Isjieru,
			})
		}
	}
	return nil
}

func (w White) ServerAddWhite(_ context.Context, request *serverApi.ServerWhiteRequest, _ *commonApi.EmptyResponse) error {
	if err := Client().ServerAddWhiteList(request.Ip, request.StrDomain); err != nil {
		return err
	}
	return nil
}

func (w White) ServerDelWhite(_ context.Context, request *serverApi.ServerWhiteRequest, _ *commonApi.EmptyResponse) error {
	if err := Client().ServerDelWhiteList(request.Ip, request.StrDomain); err != nil {
		return err
	}
	return nil
}

func (w White) ServerGetBlackList(_ context.Context, request *serverApi.ServerListRequest, response *serverApi.ServerGetBlackListResponse) error {
	if data, err := Client().ServerGetBlackList(request.Ip, request.PageNo, request.Limit); err != nil {
		return err
	} else {
		response.Total = data.Total
		response.Limit = data.Limit
		response.PageNo = data.Pageno
		response.PageCount = data.Pagecount
		response.Items = data.Items
	}
	return nil
}

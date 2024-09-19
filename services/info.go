package services

import (
	"context"
	// "git.internal.attains.cn/attains-cloud/go-api/attains/server/api"
)

type Info struct {
}

func (i Info) ServerGetBeiAnCode(_ context.Context, request *api.ServerRequest, response *api.ServerGetBeiAnCodeResponse) error {
	if data, err := Client().ServerGetBeianCode(request.Ip); err != nil {
		return err
	} else {
		response.YwId = data.Ywid
		response.YwName = data.Ywname
		response.YwType = data.Ywtype
		response.YwUserid = data.Ywuserid
		response.CrDate = data.Crdate
		response.AllowZt = data.Allowzt
		response.UseZtCount = data.Useztcount
		response.ProId = data.Proid
		for _, v := range data.Lists {
			response.List = append(response.List, &api.BeiAnCodeList{
				CodeId:    v.Codeid,
				ForYwId:   v.Forywid,
				CodeValue: v.Codevalue,
				Status:    v.Status,
				WzBsId:    v.Wzbsid,
				ZtBsId:    v.Ztbsid,
				BkDwmc:    v.BkDwmc,
				BkWzbah:   v.BkWzbah,
				BkWzym:    v.BkWzym,
				UseTime:   v.Usetime,
				YwId:      v.Ywid,
				YwName:    v.Ywname,
				YwType:    v.Ywtype,
				YwUserid:  v.Ywuserid,
				CrDate:    v.Crdate,
				AllowZt:   v.Allowzt,
			})
		}
	}
	return nil
}

func (i Info) ServerGetWebConsoleUrl(_ context.Context, request *api.ServerRequest, response *api.ServerGetWebConsoleUrlResponse) error {
	if url, err := Client().ServerGetWebConsoleUrl(request.Ip); err != nil {
		return err
	} else {
		response.Url = url
	}
	return nil
}

func (i Info) ServerOperatorLogs(_ context.Context, request *api.ServerRequest, response *api.ServerOperatorLogsResponse) error {
	if data, err := Client().ServerOperatorLogs(request.Ip); err != nil {
		return err
	} else {
		response.Total = data.Total
		response.Limit = data.Limit
		response.PageNo = data.Pageno
		response.PageCount = data.Pagecount
		for _, v := range data.Items {
			response.List = append(response.List, &api.ServerOperatorLog{
				SysId:     v.Sysid,
				TargetIp:  v.Targetip,
				TaskTime:  v.Tasktime,
				TaskState: v.Taskstate,
				Task:      v.Task,
				TaskUser:  v.Taskuser,
			})
		}
	}
	return nil
}

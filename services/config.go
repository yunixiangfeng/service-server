package services

import (
	"context"
	// commonApi "git.internal.attains.cn/attains-cloud/go-api/attains/common/api"
	// serverApi "git.internal.attains.cn/attains-cloud/go-api/attains/server/api"
)

type Config struct {
}

func (c Config) ServerGetBaseConfig(_ context.Context, _ *commonApi.EmptyRequest, response *serverApi.ServerGetBaseConfigResponse) error {
	if data, err := Client().ServerGetBaseConfig(); err != nil {
		return err
	} else {
		response.Cpu = data.Cpu
		response.Ram = data.Ram
		for _, r := range data.Room {
			ebsList := make([]*serverApi.ServerBaseConfigRoomEbs, 0)
			for _, e := range r.Ebs {
				ebsList = append(ebsList, &serverApi.ServerBaseConfigRoomEbs{
					Id:   e.Id,
					Name: e.Name,
					HCpu: e.HCpu,
				})
			}
			response.Room = append(response.Room, &serverApi.ServerBaseConfigRoom{
				Id:     r.Id,
				Name:   r.Name,
				Remark: r.Remark,
				Ebs:    ebsList,
				Ddos:   r.Ddos,
			})
		}
		for _, o := range data.Os {
			response.Os = append(response.Os, &serverApi.ServerBaseConfigOS{
				Name:    o.Name,
				Value:   o.Value,
				System:  o.System,
				Bit:     o.Bit,
				Account: o.Account,
				Port:    o.Port,
			})
		}
	}
	return nil
}

func (c Config) ServerGetOsList(_ context.Context, _ *commonApi.EmptyRequest, response *serverApi.ServerGetOsListResponse) error {
	if data, err := Client().ServerGetOsList(); err != nil {
		return err
	} else {
		for _, o := range data {
			response.List = append(response.List, &serverApi.ServerOsListNode{
				Name:  o.Name,
				Value: o.Value,
			})
		}
	}
	return nil
}

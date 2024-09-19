package setup

import (
	"context"
	// "git.internal.attains.cn/attains-cloud/service-acs/core/config/pkg"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func InitEtcd(ctx context.Context, etcd *pkg.Etcd) (*clientv3.Client, error) {
	clientV3Config := clientv3.Config{
		Endpoints:            etcd.Endpoints,
		AutoSyncInterval:     etcd.AutoSyncInterval,
		DialTimeout:          etcd.DialTimeout,
		DialKeepAliveTime:    etcd.DialKeepAliveTime,
		DialKeepAliveTimeout: etcd.DialKeepAliveTimeout,
		MaxCallSendMsgSize:   etcd.MaxCallSendMsgSize,
		MaxCallRecvMsgSize:   etcd.MaxCallRecvMsgSize,
		TLS:                  nil,
		Username:             etcd.Username,
		Password:             etcd.Password,
		RejectOldCluster:     etcd.RejectOldCluster,
		DialOptions:          nil,
		Context:              ctx,
		LogConfig:            nil,
		PermitWithoutStream:  etcd.PermitWithoutStream,
	}
	cli, err := clientv3.New(clientV3Config)
	if err != nil {
		return nil, err
	}
	return cli, nil
}

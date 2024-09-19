package core

import (
	"context"
	"fmt"
	// "git.internal.attains.cn/attains-cloud/service-acs/core/config"
	"github.com/go-redsync/redsync/v4"
	"github.com/redis/go-redis/v9"
	"github.com/sony/sonyflake"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

const DefaultKey = "default"

type PkgS struct {
	Database  map[string]*gorm.DB
	Flake     *sonyflake.Sonyflake
	Redis     map[string]redis.UniversalClient
	RedisPool map[string]*redsync.Redsync
	Logger    map[string]*zap.SugaredLogger
}

type ContainerS struct {
	Pkg  *PkgS
	Conf *config.Config
}

var Container *ContainerS

func GetConf() config.Config {
	return *Container.Conf
}

func GetDatabase(v ...interface{}) (*gorm.DB, error) {
	name := DefaultKey
	if len(v) > 0 {
		_name, _ := v[0].(string)
		if len(_name) > 0 {
			name = _name
		}
	}
	database, ok := Container.Pkg.Database[name]
	if !ok {
		return nil, fmt.Errorf("database %s not found", name)
	}
	return database, nil
}

func GetLogger(v ...interface{}) (*zap.SugaredLogger, error) {
	name := DefaultKey
	if len(v) > 0 {
		_name, _ := v[0].(string)
		if len(_name) > 0 {
			name = _name
		}
	}
	logger, ok := Container.Pkg.Logger[name]
	if !ok {
		return nil, fmt.Errorf("logger %s not found", name)
	}
	return logger, nil
}

func GetRedis(v ...interface{}) (redis.UniversalClient, error) {
	name := DefaultKey
	if len(v) > 0 {
		_name, _ := v[0].(string)
		if len(_name) > 0 {
			name = _name
		}
	}
	instance, ok := Container.Pkg.Redis[name]
	if !ok {
		return nil, fmt.Errorf("redis %s not found", name)
	}
	return instance, nil
}

func GetRedidPool(v ...interface{}) (*redsync.Redsync, error) {
	name := DefaultKey
	if len(v) > 0 {
		_name, _ := v[0].(string)
		if len(_name) > 0 {
			name = _name
		}
	}
	pool, ok := Container.Pkg.RedisPool[name]
	if !ok {
		return nil, fmt.Errorf("redis %s not found", name)
	}
	return pool, nil
}

func GetFlake() *sonyflake.Sonyflake {
	return Container.Pkg.Flake
}

func GetNextId() string {
	nextId, _ := GetFlake().NextID()
	return strconv.FormatUint(nextId, 10)
}

func GetEtcd(ctx context.Context, v ...interface{}) (*clientv3.Client, error) {
	name := DefaultKey
	if len(v) > 0 {
		_name, _ := v[0].(string)
		if len(_name) > 0 {
			name = _name
		}
	}
	var instance *clientv3.Client
	if Container.Conf.Pkg.Etcd != nil {
		for _, etcdC := range Container.Conf.Pkg.Etcd {
			if etcdC == nil {
				continue
			}
			etcdC.Name = strings.TrimSpace(etcdC.Name)
			if etcdC.Name == name {
				clientV3Config := clientv3.Config{
					Endpoints:            etcdC.Endpoints,
					AutoSyncInterval:     etcdC.AutoSyncInterval,
					DialTimeout:          etcdC.DialTimeout,
					DialKeepAliveTime:    etcdC.DialKeepAliveTime,
					DialKeepAliveTimeout: etcdC.DialKeepAliveTimeout,
					MaxCallSendMsgSize:   etcdC.MaxCallSendMsgSize,
					MaxCallRecvMsgSize:   etcdC.MaxCallRecvMsgSize,
					TLS:                  nil,
					Username:             etcdC.Username,
					Password:             etcdC.Password,
					RejectOldCluster:     etcdC.RejectOldCluster,
					DialOptions:          nil,
					Context:              ctx,
					LogConfig:            nil,
					PermitWithoutStream:  etcdC.PermitWithoutStream,
				}
				cli, err := clientv3.New(clientV3Config)
				if err != nil {
					return nil, err
				}
				instance = cli
			}
		}
	}
	if instance == nil {
		return nil, fmt.Errorf("etcd %s not found", name)
	}
	return instance, nil
}

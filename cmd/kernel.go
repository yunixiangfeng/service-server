package cmd

import (
	"context"
	"fmt"
	// "git.internal.attains.cn/attains-cloud/service-acs/core"
	// config2 "git.internal.attains.cn/attains-cloud/service-acs/core/config"
	// "git.internal.attains.cn/attains-cloud/service-acs/core/config/pkg"
	// "git.internal.attains.cn/attains-cloud/service-acs/core/metadata"
	// setup2 "git.internal.attains.cn/attains-cloud/service-acs/core/setup"
	// "git.internal.attains.cn/attains-cloud/service-acs/services"
	"github.com/go-micro/plugins/v4/registry/etcd"
	"github.com/go-micro/plugins/v4/server/grpc"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/gookit/goutil/envutil"
	"github.com/gookit/goutil/strutil"
	"github.com/joho/godotenv"
	"github.com/mitchellh/mapstructure"
	"github.com/philchia/agollo/v4"
	"github.com/redis/go-redis/v9"
	"github.com/sony/sonyflake"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/server"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
)

var defaultCfg = "cache/conf.yaml"

func readConfig() error {
	var err error
	viper.SetDefault("c", defaultCfg)
	// 设置环境变量前缀并自动加载环境变量
	viper.SetEnvPrefix("ATTAINS")
	viper.AutomaticEnv()

	if err = viper.BindPFlags(pflag.CommandLine); err != nil {
		return err
	}

	envC, configFile := viper.GetString("c"), defaultCfg
	if envC == "remote" { // 远程配置读取
		envMaps, _ := godotenv.Read("./.env")
		if envMaps == nil {
			envMaps = map[string]string{}
		}
		for k := range envMaps {
			if ev := envutil.Getenv(k); len(ev) != 0 {
				envMaps[k] = ev
			}
		}
		envutil.SetEnvMap(envMaps)
		for _, k := range metadata.RequiredEnvs {
			if ev := envutil.Getenv(k); len(ev) == 0 {
				return fmt.Errorf("you seem to have forgotten the configuration key: %s", k)
			}
		}

		namespaces := []string{"application.properties", metadata.ServiceName + ".yaml"}
		namespacesEnv := strutil.Split(envutil.Getenv("ATTAINS_APOLLO_NAMESPACES"), ",")
		if len(namespacesEnv) > 0 {
			namespaces = append(namespaces, namespacesEnv...)
		}
		err = agollo.Start(&agollo.Conf{
			AppID:              envutil.Getenv("ATTAINS_APOLLO_APP_ID"),
			Cluster:            envutil.Getenv("ATTAINS_APOLLO_CLUSTER"),
			NameSpaceNames:     namespaces,
			CacheDir:           "./cache/apollo",
			MetaAddr:           envutil.Getenv("ATTAINS_APOLLO_URL"),
			AccesskeySecret:    envutil.Getenv("ATTAINS_APOLLO_SECRET"),
			InsecureSkipVerify: false,
		})
		if err != nil {
			return err
		}
		baseConfigContent := agollo.GetContent(agollo.WithNamespace(metadata.ServiceName + ".yaml"))
		err = os.WriteFile(defaultCfg, []byte(baseConfigContent), os.ModePerm)
		if err != nil {
			return err
		}
	} else {
		configFile = envC
	}

	viper.SetConfigFile(configFile)

	return viper.ReadInConfig()
}

// Initialize 初始化
func Initialize(ctx context.Context) error {
	if err := readConfig(); err != nil {
		return err
	}
	conf := new(config2.Config)
	err := viper.Unmarshal(conf, func(d *mapstructure.DecoderConfig) {
		d.TagName = "yaml"
	})
	if err != nil {
		return err
	}

	c := new(core.ContainerS)
	c.Conf = conf
	c.Pkg = new(core.PkgS)

	development := c.Conf.App.Env == metadata.EnvDev
	debug := c.Conf.App.Debug

	// initialize logger
	for _, l := range c.Conf.Pkg.Logger {
		logger, err := setup2.InitLogger(ctx, development, debug)
		if err != nil {
			return fmt.Errorf("init logger err: %v", err)
		}
		if c.Pkg.Logger == nil {
			c.Pkg.Logger = map[string]*zap.SugaredLogger{}
		}
		c.Pkg.Logger[l.Name] = logger.Sugar()
	}

	// initialize database
	if c.Conf.Pkg.Database != nil {
		for _, database := range c.Conf.Pkg.Database {
			db, err := setup2.InitDatabase(ctx, database)
			if err != nil {
				return fmt.Errorf("init database err: %v", err)
			}
			if c.Pkg.Database == nil {
				c.Pkg.Database = map[string]*gorm.DB{}
			}
			if debug {
				db = db.Debug()
			}
			c.Pkg.Database[database.Name] = db
		}
	}

	// initialize redis
	if c.Conf.Pkg.Redis != nil {
		for _, r := range c.Conf.Pkg.Redis {
			redisClient, err := setup2.InitRedis(ctx, r)
			if err != nil {
				return fmt.Errorf("init redis err: %v", err)
			}
			if c.Pkg.Redis == nil {
				c.Pkg.Redis = map[string]redis.UniversalClient{}
				c.Pkg.RedisPool = map[string]*redsync.Redsync{}
			}
			c.Pkg.Redis[r.Name] = redisClient
			c.Pkg.RedisPool[r.Name] = redsync.New(goredis.NewPool(redisClient))
		}
	}

	// initialize snowflake
	sonyFlakeSettings := sonyflake.Settings{}
	c.Pkg.Flake = sonyflake.NewSonyflake(sonyFlakeSettings)

	core.Container = c

	return nil
}

func Run(ctx context.Context) error {
	if err := Initialize(ctx); err != nil {
		return err
	}
	c := core.Container
	serverC := c.Conf.App.Server
	grpcServer := grpc.NewServer(
		server.Name(metadata.ServiceName),
		server.Address(fmt.Sprintf(":%d", serverC.SocketPort)),
		server.Advertise(fmt.Sprintf("%s:%d", serverC.SocketAddress, serverC.SocketPort)),
	)

	if err := services.RegisterServices(grpcServer); err != nil {
		return err
	}

	opts := make([]micro.Option, 0)
	opts = append(opts,
		micro.Context(ctx),
		micro.Name(metadata.ServiceName),
		micro.Server(grpcServer),
	)

	etcdCfgMaps := make(map[string]*pkg.Etcd)
	if etcdList := c.Conf.Pkg.Etcd; etcdList != nil {
		for _, e := range etcdList {
			etcdCfgMaps[e.Name] = e
		}
	}
	if registries := c.Conf.App.Registries; registries != nil {
		etcdAddrs := make([]string, 0)
		for _, r := range registries {
			if r.Type == "etcd" {
				etcdCfg, ok := etcdCfgMaps[r.Connection]
				if !ok {
					continue
				}
				etcdAddrs = append(etcdAddrs, etcdCfg.Endpoints...)
			}
		}
		if len(etcdAddrs) > 0 {
			opts = append(opts, micro.Registry(
				etcd.NewRegistry(
					registry.Addrs(etcdAddrs...),
				),
			))
		}
	}

	service := micro.NewService(opts...)
	service.Init()
	return service.Run()
}

package services

import (
	// "git.internal.attains.cn/attains-cloud/cloud-provider-sdk-west/auth"
	// "git.internal.attains.cn/attains-cloud/cloud-provider-sdk-west/business"
	// "git.internal.attains.cn/attains-cloud/cloud-provider-sdk-west/model"
	// "git.internal.attains.cn/attains-cloud/go-api/attains/server/api"
	"github.com/gookit/goutil/envutil"
	"go-micro.dev/v4/server"
)

func RegisterServices(grpcServer server.Server) error {
	var err error
	if err = api.RegisterServerBaseHandler(grpcServer, new(Base)); err != nil {
		return err
	}
	if err = api.RegisterServerConfigHandler(grpcServer, new(Config)); err != nil {
		return err
	}
	if err = api.RegisterServerExpansionHandler(grpcServer, new(Expansion)); err != nil {
		return err
	}
	if err = api.RegisterServerFirewallHandler(grpcServer, new(Firewall)); err != nil {
		return err
	}
	if err = api.RegisterServerFirewallRuleHandler(grpcServer, new(FirewallRule)); err != nil {
		return err
	}
	if err = api.RegisterServerFormalHandler(grpcServer, new(Formal)); err != nil {
		return err
	}
	if err = api.RegisterServerHostSnapHandler(grpcServer, new(HostSnap)); err != nil {
		return err
	}
	if err = api.RegisterServerHostSnapMountHandler(grpcServer, new(HostSnapMount)); err != nil {
		return err
	}
	if err = api.RegisterServerHostSnapSetDiyHandler(grpcServer, new(HostSnapSetDiy)); err != nil {
		return err
	}
	if err = api.RegisterServerInfoHandler(grpcServer, new(Info)); err != nil {
		return err
	}
	if err = api.RegisterServerIpHandler(grpcServer, new(Ip)); err != nil {
		return err
	}
	if err = api.RegisterServerManageHandler(grpcServer, new(Manage)); err != nil {
		return err
	}
	if err = api.RegisterServerRenewHandler(grpcServer, new(Renew)); err != nil {
		return err
	}
	if err = api.RegisterServerSecurityInvasionHandler(grpcServer, new(SecurityInvasion)); err != nil {
		return err
	}
	if err = api.RegisterServerStatusHandler(grpcServer, new(Status)); err != nil {
		return err
	}
	if err = api.RegisterServerUpgradeHandler(grpcServer, new(Upgrade)); err != nil {
		return err
	}
	if err = api.RegisterServerWhiteHandler(grpcServer, new(White)); err != nil {
		return err
	}
	return nil
}

func Client() *business.Client {
	return business.New(auth.New(&model.Auth{
		Username:    envutil.Getenv("CLIENT_USERNAME"),
		ApiPassword: envutil.Getenv("CLIENT_PASSWORD"),
	}))
}

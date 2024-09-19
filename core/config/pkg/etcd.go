package pkg

import "time"

type Etcd struct {
	Name                 string        `json:"name" xml:"name" yaml:"name"`
	Endpoints            []string      `json:"endpoints" xml:"endpoints" yaml:"endpoints"`
	AutoSyncInterval     time.Duration `json:"auto_sync_interval" xml:"auto_sync_interval" yaml:"auto_sync_interval"`
	DialTimeout          time.Duration `json:"dial_timeout" xml:"dial_timeout" yaml:"dial_timeout"`
	DialKeepAliveTime    time.Duration `json:"dial_keep_alive_time" xml:"dial_keep_alive_time" yaml:"dial_keep_alive_time"`
	DialKeepAliveTimeout time.Duration `json:"dial_keep_alive_timeout" xml:"dial_keep_alive_timeout" yaml:"dial_keep_alive_timeout"`
	MaxCallSendMsgSize   int           `json:"max_call_send_msg_size" xml:"max_call_send_msg_size" yaml:"max_call_send_msg_size"`
	MaxCallRecvMsgSize   int           `json:"max_call_recv_msg_size" xml:"max_call_recv_msg_size" yaml:"max_call_recv_msg_size"`
	Username             string        `json:"username" xml:"username" yaml:"username"`
	Password             string        `json:"password" xml:"password" yaml:"password"`
	RejectOldCluster     bool          `json:"reject_old_cluster" xml:"reject_old_cluster" yaml:"reject_old_cluster"`
	PermitWithoutStream  bool          `json:"permit_without_stream" xml:"permit_without_stream" yaml:"permit_without_stream"`
}

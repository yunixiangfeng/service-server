package pkg

type Redis struct {
	Name             string   `json:"name" xml:"name" yaml:"name"`
	Addrs            []string `json:"addrs" xml:"addrs" yaml:"addrs"`
	DB               int      `json:"db" xml:"db" yaml:"db"`
	Username         string   `json:"username" xml:"username" yaml:"username"`
	Password         string   `json:"password" xml:"password" yaml:"password"`
	SentinelPassword string   `json:"sentinel_password" xml:"sentinel_password" yaml:"sentinel_password"`
	MaxRetries       int      `json:"max_retries" xml:"max_retries" yaml:"max_retries"`
	PoolFIFO         bool     `json:"pool_fifo" xml:"pool_fifo" yaml:"pool_fifo"`
	PoolSize         int      `json:"pool_size" xml:"pool_size" yaml:"pool_size"`
	MinIdleConns     int      `json:"min_idle_conns" xml:"min_idle_conns" yaml:"min_idle_conns"`
	MaxRedirects     int      `json:"max_redirects" xml:"max_redirects" yaml:"max_redirects"`
	ReadOnly         bool     `json:"readonly" xml:"readonly" yaml:"readonly"`
	RouteByLatency   bool     `json:"route_by_latency" xml:"route_by_latency" yaml:"route_by_latency"`
	RouteRandomly    bool     `json:"route_randomly" xml:"route_randomly" yaml:"route_randomly"`
	MasterName       string   `json:"master_name" xml:"master_name" yaml:"master_name"`
	Check            bool     `json:"check" xml:"check" yaml:"check"`
}

package app

type Server struct {
	SocketAddress string            `json:"socket_address" xml:"socket_address" yaml:"socket_address"`
	SocketPort    uint16            `json:"socket_port" xml:"socket_port" yaml:"socket_port"`
}

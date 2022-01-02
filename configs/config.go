package configs

import "fmt"

const (
	ChainAddress = "54.211.69.39"
	// ChainAddress = "ec2-3-80-211-46.compute-1.amazonaws.com"
	ChainRPCPort  = 26657
	ChainRESTPort = 1317
	ChainWSPath   = "/websocket"
	ChainName     = "saage-testnet-1"
)

func WebsocketAddr() string {
	return fmt.Sprintf("tcp://%s:%d", ChainAddress, ChainRPCPort)
}

func HttpAddr() string {
	return fmt.Sprintf("http://%s:%d", ChainAddress, ChainRESTPort)
}

package subscriptions

import (
	"log"

	"github.com/cosmos-listener/configs"

	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
)

var Client *rpchttp.HTTP

func StartWebSocket() (client *rpchttp.HTTP, err error) {
	socketURL := configs.WebsocketAddr()
	log.Println(socketURL)
	//socketURL := configs.ChainAddress

	Client, err = rpchttp.New(socketURL, configs.ChainWSPath)

	if err != nil {
		log.Println("Failed to create websocket client!!")
	}

	if err = Client.Start(); err != nil {
		log.Println("Failed to start websocket client!!", err)
	} else {
		log.Println("Successfully started Websocket client")
	}

	NewBlockSub()
	return client, err
}

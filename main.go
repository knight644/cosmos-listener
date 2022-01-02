package main

import (
	"github.com/cosmos-listener/record"
	"github.com/cosmos-listener/subscriptions"
)

func Run() {
	record.CreateMonitorCSVWriter()
	subscriptions.StartWebSocket()
}

func main() {
	Run()
}

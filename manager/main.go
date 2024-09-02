package main

import (
	"fmt"
	"publish-expcetion/manager/nats"
	sub "publish-expcetion/manager/subscription"
)

var subs sub.Subscription = nats.DefaultNatsSubscription{}

func main() {
	fmt.Println("Iniciando Publish Exception Manager")
	subs.Reader()
}

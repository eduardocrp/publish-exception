package main

import (
	"flag"
	"log"
	"publish-expcetion/publisher/nats"
	pub "publish-expcetion/publisher/publisher"
)

var publisher = nats.DefaultNatsPublisher{}

func main() {
	message := loadMessageArgs()

	if err := message.Publish(publisher); err != nil {
		log.Fatalf("Falha ao publicar: %s", err.Error())
	}
}

func loadMessageArgs() pub.MessageException {
	appName := flag.String("app_name", "", "Aplicação de origem")
	criticalLevel := flag.String("critical_level", "", "Nível de exceção")
	id := flag.String("id", "", "Identificador")
	flow := flag.String("flow", "", "Fluxo a ser mapeado")
	msg := flag.String("msg", "", "Messagem de exceção")
	flag.Parse()

	var message = pub.MessageException{
		ApplicationName: *appName,
		CriticalLevel:   *criticalLevel,
		ID:              *id,
		Flow:            *flow,
		Message:         *msg,
	}
	return message
}

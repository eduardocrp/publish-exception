package main

import (
	"flag"
	"log"
	pub "publish-expcetion/publisher/publishers"
	"publish-expcetion/publisher/publishers/nats"
	"publish-expcetion/publisher/publishers/redis"
)

var publisher pub.Publisher

func main() {
	loadPublisher()

	message := loadMessageArgs()

	if err := message.Publish(publisher); err != nil {
		log.Fatalf("Falha ao publicar: %s", err.Error())
	}
}

func loadPublisher() {
	messageProvider := flag.String("message_provider", "default", "Tipo de provedor de mensageria (redis ou nats)")

	switch *messageProvider {
	case "redis", "default":
		publisher = redis.DefaultRedisPublisher{}
	case "nats":
		publisher = nats.DefaultNatsPublisher{}
	default:
		log.Fatalln("Não selecionado provedor de mensageria compatível. Provedores disponíveis: redis ou nats")
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

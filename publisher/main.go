package main

import (
	"flag"
	"log"
	pub "publish-expcetion/publisher/publisher"
	"publish-expcetion/publisher/publisher/kafka"
)

var publisher = kafka.DefaultKafkaPublisher{}

func main() {
	appName := flag.String("app_name", "origem_app", "Aplicação de origem")
	criticalLevel := flag.String("critical_level", "ERROR", "Nível de exceção")
	flag.Parse()

	var message = pub.MessageException{
		ApplicationName: *appName,
		CriticalLevel:   *criticalLevel,
	}

	message.Flow = "flow-test"
	if err := message.Publish(publisher); err != nil {
		log.Fatalf("Falha ao publicar: %s", err.Error())
	}
}

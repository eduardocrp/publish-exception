package nats

import (
	"fmt"
	"log"
	"sync"

	"github.com/nats-io/nats.go"
)

type DefaultNatsSubscrition struct{}

const (
	streamName        = "publish-exception"
	durablePrefixName = "publish-exception-manager"
)

func (DefaultNatsSubscrition) Reader() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	js, err := nc.JetStream()
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup

	for _, sub := range loadSubs(js, streamName) {
		wg.Add(1)
		go func(sub string) {
			defer wg.Done()
			subscribe(js, sub)
		}(sub)
	}

	wg.Wait()
}

func loadSubs(js nats.JetStreamContext, streamName string) []string {
	subs, err := js.StreamInfo(streamName)
	if err != nil {
		log.Fatal("Erro ao listar assuntos do stream ", streamName, " \nErro: ", err)
	}
	log.Println("Se inscrevendo nos assuntos:")
	for _, sub := range subs.Config.Subjects {
		log.Println(sub)
	}
	return subs.Config.Subjects
}

func subscribe(js nats.JetStreamContext, sub string) {
	_, err := js.Subscribe(sub, func(msg *nats.Msg) {
		log.Printf("Assunto: %s\nmensagem: %s\n", sub, string(msg.Data))
		msg.Ack()
	}, nats.Durable(fmt.Sprintf("%s_%s", durablePrefixName, sub)), nats.ManualAck())
	if err != nil {
		log.Println("Falha ao se increver em assunto ", sub, " \nErro: ", err)
	} else {
		log.Println("Incrito em assunto " + sub)
		select {}
	}
}

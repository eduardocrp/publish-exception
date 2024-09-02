package nats

import (
	"encoding/json"
	"fmt"
	"log"
	"publish-expcetion/manager/connectors"
	"sync"

	"github.com/nats-io/nats.go"
)

type DefaultNatsSubscription struct {
	kvManager nats.KeyValueManager
}

const (
	streamName                 = "publish-exception"
	durablePrefixName          = "publish-exception-manager"
	connectorsConfigBucketName = "publish-exception-manager"
	connectorsConfigPrefixName = "publish-exception-connectors-config"
)

func (d DefaultNatsSubscription) Reader() {
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

	for _, sub := range d.loadSubs(js, streamName) {
		wg.Add(1)
		go func(sub string) {
			defer wg.Done()
			d.subscribe(js, sub)
		}(sub)
	}

	wg.Wait()
}

func (DefaultNatsSubscription) loadSubs(js nats.JetStreamContext, streamName string) []string {
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

func (d DefaultNatsSubscription) subscribe(js nats.JetStreamContext, sub string) {
	_, err := js.Subscribe(sub, func(msg *nats.Msg) {
		message := string(msg.Data)
		log.Printf("Assunto: %s\nmensagem: %s\n", sub, message)
		go d.sendToConectors(sub, message)
		msg.Ack()
	}, nats.Durable(fmt.Sprintf("%s_%s", durablePrefixName, sub)), nats.ManualAck())
	if err != nil {
		log.Println("Falha ao se increver em assunto ", sub, " \nErro: ", err)
	} else {
		log.Println("Incrito em assunto " + sub)
		select {}
	}
}

func (d DefaultNatsSubscription) sendToConectors(sub string, message string) {
	// Publicar cada conector em go routines diferentes
	connectors, err := d.getConnectorConfig(connectorsConfigPrefixName + "_" + sub)
	if err == nil && connectors != nil {
		for _, con := range *connectors {
			go con.SendMessage(message) //TODO: Incluir context timeout
		}
	}
}

func (d DefaultNatsSubscription) getConnectorConfig(subKey string) (*[]connectors.ConnectorConfig, error) {
	kv, err := d.kvManager.KeyValue(subKey)
	if err != nil {
		log.Println("Falha ao se conectar ao KV bucket ", connectorsConfigBucketName, " \nErro: ", err)
	}

	entry, err := kv.Get(subKey)
	if err != nil {
		log.Println("Falha ao coletar a configuração ", subKey, " \nErro: ", err)
	}

	var connectorConfig *[]connectors.ConnectorConfig

	err = json.Unmarshal(entry.Value(), connectorConfig)
	if err != nil {
		log.Println("Falha ao deserializar a configuração ", entry.Key(), " \nErro: ", err)
		return nil, err
	}

	return connectorConfig, nil
}

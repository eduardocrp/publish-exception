package redis

import (
	"context"
	"encoding/json"
	"fmt"
	pub "publish-expcetion/publisher/publishers"
	"time"

	"github.com/redis/go-redis/v9"
)

type DefaultRedisPublisher struct{}

func (DefaultRedisPublisher) Publish(msg *pub.MessageException) error {
	var ctx = context.Background()

	// Configurando o cliente Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Endereço do Redis
		Password: "",               // Senha se houver (se não houver, deixar vazio)
		DB:       0,                // Usar o banco de dados Redis padrão (0)
	})

	// Mensagem que será publicada na fila
	message, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	// Nome da fila (lista) onde a mensagem será publicada
	queue := msg.ApplicationName + "-publish-exception"

	// Tempo para expirar por padrão definido para 600 seg (10 min)
	err = rdb.Expire(ctx, queue, 600*time.Second).Err()
	if err == nil {

		// Publicando a mensagem na fila
		err = rdb.LPush(ctx, queue, message).Err()
		if err == nil {
			fmt.Printf("Mensagem publicada com sucesso na fila '%s'\n", queue)
		}

	}

	return err
}

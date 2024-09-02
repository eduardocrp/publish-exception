package connectors

type Message struct {
	Content string `json:"text"`
}

type ConnectorConfig struct {
	Url string
}

type Connector interface {
	SendMessage(message string) error
}

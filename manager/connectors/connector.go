package connectors

type Message struct {
	Content string `json:"text"`
}

type Connector interface {
	SendMessage(url string, message string) error
}

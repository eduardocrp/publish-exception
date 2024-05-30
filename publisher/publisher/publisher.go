package publisher

type MessageException struct {
	ApplicationName string
	CriticalLevel   string
	ID              string
	Flow            string
	Message         string
}

type Publisher interface {
	Publish(msg *MessageException) error
}

func (msg *MessageException) Publish(p Publisher) error {
	return p.Publish(msg)
}

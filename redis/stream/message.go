package stream

type Message struct {
	ID     string
	Stream string
	Values map[string]interface{}
}

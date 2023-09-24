package websocket

type ClientObserver interface {
	Notify(client *Client)
}

type ClientListener struct {
	observers []ClientObserver
}

func NewClientListener() *ClientListener {
	return &ClientListener{}
}

var listener *ClientListener

func GetClientListener() *ClientListener {
	if listener == nil {
		listener = NewClientListener()
	}
	return listener
}

func (l *ClientListener) Register(observer ClientObserver) *ClientListener {
	l.observers = append(l.observers, observer)
	return l
}

func (l *ClientListener) Notify(client *Client) {
	for _, observer := range l.observers {
		observer.Notify(client)
	}
}

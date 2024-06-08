package eventhandler

type Event interface {
	GetID() string
	GetType() string
	GetBody() string
}

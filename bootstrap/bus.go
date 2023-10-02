package bootstrap

import "errors"

type Module struct {
	Name     string
	Handlers map[string]func(interface{}) interface{}
	Messages chan Message
}

type Message struct {
	Action   string
	Data     interface{}
	Error    error
	Response chan Message
}

func NewBus(name string, messages chan Message) *Module {
	return &Module{
		Name:     name,
		Handlers: make(map[string]func(interface{}) interface{}),
		Messages: messages,
	}
}

func (m *Module) AddHandler(action string, handler func(interface{}) interface{}) {
	m.Handlers[action] = handler
}

func (m *Module) Start() {
	for {
		select {
		case msg := <-m.Messages:
			if handler, exists := m.Handlers[msg.Action]; exists {
				responseData := handler(msg.Data)
				msg.Response <- Message{Action: msg.Action, Data: responseData}
			} else {
				msg.Response <- Message{Action: "error", Error: errors.New("no handler for action")}
			}
		}
	}
}

func InitBusComm() (*Module, chan Message) {
	messages := make(chan Message)

	bus := NewBus("bus", messages)

	go bus.Start()

	return bus, bus.Messages
}

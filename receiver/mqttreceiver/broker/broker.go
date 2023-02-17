package broker

import (
	"github.com/fhmq/hmq/broker"
)

type mqttBroker struct {
	broker *broker.Broker
}

func New(cfg *Config) (*mqttBroker, error) {
	b, err := broker.NewBroker(cfg.Config)
	if err != nil {
		return nil, err
	}

	return &mqttBroker{
		b,
	}, nil
}

func (b *mqttBroker) Start() {
	b.broker.Start()
}

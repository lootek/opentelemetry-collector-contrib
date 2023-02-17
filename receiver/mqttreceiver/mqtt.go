package mqttreceiver

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
)

type Broker interface {
	Start()
}

type Message interface {
	Topic() string
	Data() interface{}
}

type Subscriber interface {
	Subscribe() (<-chan Message, error)
}

type mqttCollector struct {
	uri   *url.URL
	topic string
}

func (m *mqttCollector) Start(ctx context.Context, host component.Host) error {
	panic("implement me")
}

func (m *mqttCollector) Shutdown(ctx context.Context) error {
	panic("implement me")
}

func New(params component.ReceiverCreateSettings, c Config, metrics consumer.Metrics) (component.MetricsReceiver, error) {
	// mqtt.ERROR = logger
	// mqtt.CRITICAL = logger
	// mqtt.WARN = logger
	// mqtt.DEBUG = logger

	url, err := url.Parse(cfg.Endpoint)
	if err != nil {
		return nil, err
	}

	return &mqttCollector{
		uri:   url,
		topic: cfg.Topic,
	}, nil
}

type message struct {
	topic string
	data  string
}

func (m message) Topic() string {
	return strings.ReplaceAll(m.topic, "$SYS/", "")
}

func (m message) Data() interface{} {
	return m.data
}

func (m *mqttCollector) Subscribe() (<-chan Message, error) {
	dataCh := make(chan Message)

	client, err := connect("snap-plugin-collector-mq", m.uri)
	if err != nil {
		return nil, err
	}

	go func() {
		client.Subscribe(m.topic, 0, func(client mqtt.Client, msg mqtt.Message) {
			data := string(msg.Payload())
			fmt.Printf("* [%s] %s\n", msg.Topic(), data)
			dataCh <- message{
				topic: msg.Topic(),
				data:  data,
			}
		})
	}()

	return dataCh, nil
}

func connect(clientId string, uri *url.URL) (mqtt.Client, error) {
	opts := createClientOptions(clientId, uri)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(time.Second) {
	}
	if err := token.Error(); err != nil {
		return nil, err
	}

	return client, nil
}

func createClientOptions(clientId string, uri *url.URL) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", uri.Host))
	opts.SetUsername(uri.User.Username())
	password, _ := uri.User.Password()
	opts.SetPassword(password)
	opts.SetClientID(clientId)

	return opts
}

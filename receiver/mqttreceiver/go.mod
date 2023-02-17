module github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mqttreceiver

go 1.17

require (
	github.com/stretchr/testify v1.7.0
	go.opentelemetry.io/collector v0.38.1-0.20211103011348-c24dfeb047a8
	go.opentelemetry.io/collector/model v0.38.1-0.20211103011348-c24dfeb047a8 // indirect
)

require (
	github.com/eclipse/paho.mqtt.golang v1.3.5
	github.com/fhmq/hmq v0.0.0-20210810024638-1d6979189a22
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal => ../../internal/coreinternal

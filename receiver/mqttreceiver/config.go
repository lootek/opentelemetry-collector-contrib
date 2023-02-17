// Copyright 2020, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mqttreceiver

import (
	"time"

	"go.opentelemetry.io/collector/config"
)

// Config defines configuration for mqtt receiver.
type Config struct {
	config.ReceiverSettings `mapstructure:",squash"`
	Endpoint                string        `mapstructure:"endpoint"`
	Topic                   string        `mapstructure:"topic"`
	AggregationInterval     time.Duration `mapstructure:"aggregation_interval"`
	EnableMetricType        bool          `mapstructure:"enable_metric_type"`
	IsMonotonicCounter      bool          `mapstructure:"is_monotonic_counter"`
}

func DefaultConfig() *Config {
	return &Config{
		Endpoint: "tcp://192.168.10.18:1883",
		Topic:    "#",
	}
}

func (c *Config) validate() error {
	return nil
}

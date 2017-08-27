package main

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

// Config hold all service configuration variables.
type Config struct {
	ServiceName       string `envconfig:"SERVICE_NAME"        default:"event_emitter"`
	ServiceInstanceID string `envconfig:"SERVICE_INSTANCE_ID" default:"1"`

	LogLevel  string `envconfig:"LOG_LEVEL" default:"info" desc:"Zap log level string"`
	SentryDSN string `envconfig:"SENTRY_DSN"`

	STANURL                 string        `envconfig:"STAN_URL"                   default:"nats://localhost:4222"`
	STANClusterID           string        `envconfig:"STAN_CLUSTER_ID"            default:"steemwatch"`
	STANClientID            string        `envconfig:"STAN_CLIENT_ID"             default:"block_processor"`
	STANConnectWait         time.Duration `envconfig:"STAN_CONNECT_WAIT"          default:"2s"`
	STANPubAckWait          time.Duration `envconfig:"STAN_PUB_ACK_WAIT"          default:"30s"`
	STANInputSubject        string        `envconfig:"STAN_INPUT_SUBJECT"         default:"block_operations"`
	STANOutputSubjectPrefix string        `envconfig:"STAN_OUTPUT_SUBJECT_PREFIX" default:"events"`
}

// LoadConfig loads service configuration from the environment.
func LoadConfig() (*Config, error) {
	var config Config
	if err := envconfig.Process("", &config); err != nil {
		return nil, err
	}
	return &config, nil
}

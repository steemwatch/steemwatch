package main

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

// Config hold all service configuration variables.
type Config struct {
	ServiceName       string `envconfig:"SERVICE_NAME"        default:"block_processor"`
	ServiceInstanceID string `envconfig:"SERVICE_INSTANCE_ID" default:"1"`

	LogLevel  string `envconfig:"LOG_LEVEL" default:"info" desc:"Zap log level string"`
	SentryDSN string `envconfig:"SENTRY_DSN"`

	SteemdRPCEndpointAddress string        `envconfig:"STEEMD_RPC_ENDPOINT_ADDRESS" required:"true"`
	SteemdMaxReconnectDelay  time.Duration `envconfig:"STEEMD_MAX_RECONNECT_DELAY"  default:"1m"`
	SteemdDialTimeout        time.Duration `envconfig:"STEEMD_DIAL_TIMEOUT"         default:"2s"`
	SteemdWriteTimeout       time.Duration `envconfig:"STEEMD_WRITE_TIMEOUT"        default:"6s"`
	SteemdReadTimeout        time.Duration `envconfig:"STEEMD_READ_TIMEOUT"         default:"10s"`

	STANURL           string        `envconfig:"STAN_URL"            default:"nats://localhost:4222"`
	STANClusterID     string        `envconfig:"STAN_CLUSTER_ID"     default:"steemwatch"`
	STANClientID      string        `envconfig:"STAN_CLIENT_ID"      default:"block_processor"`
	STANConnectWait   time.Duration `envconfig:"STAN_CONNECT_WAIT"   default:"2s"`
	STANPubAckWait    time.Duration `envconfig:"STAN_PUB_ACK_WAIT"   default:"30s"`
	STANOutputSubject string        `envconfig:"STAN_OUTPUT_SUBJECT" default:"steem.block_operations"`

	STANStateSubject string `envconfig:"STAN_STATE_SUBJECT" default:"_state.block_processor"`
}

// LoadConfig loads service configuration from the environment.
func LoadConfig() (*Config, error) {
	var config Config
	if err := envconfig.Process("", &config); err != nil {
		return nil, err
	}
	return &config, nil
}

package mqtt

import (
	paho "github.com/eclipse/paho.mqtt.golang"
)

// ClientConfig is the configuration for the MQTT Client struct that were provided by Eclipse
type ClientConfig struct {
	Broker         string
	ClientID       string
	MessageHandler paho.MessageHandler
}

// NewMQTTClient initialize the options through the config that were passed down and returns a MQTT client
func NewMQTTClient(config ClientConfig) paho.Client {
	opts := paho.NewClientOptions()
	opts.AddBroker(config.Broker)
	opts.SetClientID(config.ClientID)
	opts.SetDefaultPublishHandler(config.MessageHandler)
	return paho.NewClient(opts)
}

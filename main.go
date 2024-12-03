package main

import (
	"bc-public-service-assessment/mqtt"
	"bc-public-service-assessment/utils"
	"flag"
	"fmt"
	"log"
	"time"

	pahoMqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found. Using default values.")
	}

	broker := utils.GetEnv("MQTT_BROKER", "tcp://test.mosquitto.org:1883")
	clientID := utils.GetEnv("MQTT_CLIENT_ID", "go-mqtt-client-123")
	baseInputTopic := utils.GetEnv("MQTT_BASE_INPUT_TOPIC", "BRE/calculateWinterSupplementInput")

	customTopicID := flag.String("topicID", "", "Custom topic ID for MQTT input subscription")
	flag.Parse()

	if *customTopicID == "" {
		log.Fatalf("Error: --topicID flag is required")
	}

	inputTopic := fmt.Sprintf("%s/%s", baseInputTopic, *customTopicID)

	mqttConfig := mqtt.ClientConfig{
		Broker:   broker,
		ClientID: clientID,
		MessageHandler: func(client pahoMqtt.Client, message pahoMqtt.Message) {
			mqtt.HandleMessage(client, message, *customTopicID)
		},
	}

	client := mqtt.NewMQTTClient(mqttConfig)

	// Connect to the MQTT broker
	for {
		if token := client.Connect(); token.Wait() && token.Error() != nil {
			log.Printf("Connection error, retrying: %s", token.Error())
			time.Sleep(5 * time.Second)
			continue
		}
		break
	}
	log.Println("Connected to MQTT broker")

	if token := client.Subscribe(inputTopic, 0, nil); token.Wait() && token.Error() != nil {
		log.Fatalf("Error subscribing to topic: %s\n", token.Error())
	}
	log.Printf("Subscribed to topic: %s\n", inputTopic)

	// Keep program running
	select {}
}

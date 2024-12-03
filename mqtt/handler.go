package mqtt

import (
	"bc-public-service-assessment/utils"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"bc-public-service-assessment/models"
	"bc-public-service-assessment/rules"

	paho "github.com/eclipse/paho.mqtt.golang"
)

// HandleMessage is responsible for processing incoming MQTT messages, validating their payload,
// calculating the supplement using business rules, and publishing the result back to a specified output topic.
func HandleMessage(client paho.Client, msg paho.Message, topicID string) {
	log.Printf("Message received on topic: %s\n", msg.Topic())
	log.Printf("Payload: %s\n", string(msg.Payload()))

	var input models.InputData
	err := json.Unmarshal(msg.Payload(), &input)
	if err != nil {
		log.Printf("Error decoding JSON: %s\n", err)
		return
	}

	if err := ValidateInput(input); err != nil {
		log.Printf("Validation error: %s\n", err)
		return
	}

	output := rules.CalculateSupplement(input)

	baseOutputTopic := utils.GetEnv("MQTT_BASE_OUTPUT_TOPIC", "BRE/calculateWinterSupplementOutput")
	outputTopic := fmt.Sprintf("%s/%s", baseOutputTopic, topicID)
	if err := publishOutput(client, outputTopic, output); err != nil {
		log.Printf("Error publishing output: %s\n", err)
	}
}

// ValidateInput takes the input from the json input that was parsed the InputData struct and validates it
func ValidateInput(input models.InputData) error {
	if input.ID == "" {
		return errors.New("missing ID")
	}
	if input.FamilyComposition != "single" && input.FamilyComposition != "couple" {
		return fmt.Errorf("invalid family composition: %s", input.FamilyComposition)
	}
	if input.NumberOfChildren < 0 {
		return errors.New("number of children cannot be negative")
	}
	return nil
}

// publishOutput is responsible for publishing a message from a configured client on a specified topic
func publishOutput(client paho.Client, topic string, output models.OutputData) error {
	outputJSON, err := json.Marshal(output)
	if err != nil {
		return fmt.Errorf("error encoding JSON: %w", err)
	}

	token := client.Publish(topic, 0, false, outputJSON)
	token.Wait()
	if token.Error() != nil {
		return fmt.Errorf("error publishing message: %w", token.Error())
	}

	log.Printf("Published message to topic %s: %s\n", topic, string(outputJSON))
	return nil
}

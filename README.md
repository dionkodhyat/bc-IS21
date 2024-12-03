# Winter Suplement Calculator

A simple application for publishing and subscribing to MQTT topics. This project processes input data and calculates supplements based on a rule engine.


## Features

- Connect to an MQTT broker for publishing and subscribing
- Process incoming messages to calculate winter supplement amount using business logic
- Handles validation for input data
- Publishes result to an output MQTT topic


## Instructions

Due to Golang being a compiled language, the executable binaries file have been build and stored in the `/bin` directory. The following
steps to run


1. Clone the Repository
```git
git clone https://github.com/dionkodhyat/bc-IS21.git
cd bc-IS21
cd bin
```
---

2. Environment Variables (optional):

Create a `.env` file in the root directory to set up environment variables:
```git
MQTT_BROKER=tcp://test.mosquitto.org:1883
MQTT_CLIENT_ID=go-mqtt-client
MQTT_INPUT_TOPIC=BRE/calculateWinterSupplementInput
MQTT_OUTPUT_TOPIC=BRE/calculateWinterSupplementOutput
```
**Note:** default values have been configured if environment variables are not provided

---
3. Execute binaries through the following command
```git
 ./winter_supplement_calculator --topicID=<topic-id>
```
With the `<topic-id>` value be the specified one to be listened to


## Run the program (optional)
This section is for anyone who want to work and modify the code

### Prerequisites

Before setting up the project, ensure you have the following installed on your system:
- [Go](https://go.dev/doc/install) (optional, v1.19 or later)
- [Mosquitto](https://mosquitto.org/download/) (optional, for testing through the command line)


---


1. Ensure the Go programming language has been installed. See Prerequsites 

---


2. Install Dependencies:
```git
go mod tidy
```


---


3. Run the application
Run the application by specifying the topicID via the command-line arguments. See step 1 of the previous section if .env are provided
```git
go run main.go --topicID=<your-topic-id>
```
Replace `<your-topic-id>` with the desired topic ID for the MQTT input and output topic, for example:
```git
go run main.go --topicID=4ee00a09-f375-4d76-89f0-56fa79bb6e92
```


---


## Testing Input and Output with Mosquitto
As of the submission of this assignment, I was unable to retrieve payloads from the web app. However, I verified that the subscriber and publisher logic works using the Mosquitto CLI tool.

The Mosquitto CLI tools are utilities provided with the Mosquitto MQTT broker that allow users to interact with MQTT topics directly from the terminal. To verify the broker and subscriber functionality, follow the steps below:
1. Open a terminal and subscribe to a topic
```git
mosquitto_sub -h test.mosquitto.org -p 1883 -t "BRE/calculateWinterSupplementInput/<topic-id>"
```
Replace <topic-id> with the same topic ID passed to the program.

2. Open another terminal tab and subscribe to the output topic. Ensure the
```git
mosquitto_sub -h test.mosquitto.org -p 1883 -t "BRE/calculateWinterSupplementOutput/<topic-id>"
```
Again, replace <topic-id> with the same topic ID passed to the program.

3. Open a third terminal and publish a message to the input topic
```git
mosquitto_pub -h test.mosquitto.org -p 1883 -t "BRE/calculateWinterSupplementInput/<topic-id>" -m '{"id": "37632486-e85a-4251-b82c-ab308ad9491b", "numberOfChildren": 2, "familyComposition": "single", "familyUnitInPayForDecember": true}'
```
You can modify the payload values as needed, and ensure the topic-id matches the first two steps.

---
### Verify the Workflow


1.  In the terminal subscribed to the input topic, you should see the message you just published.
2.	Processing: If the program is running, it will process the message using the business logic. The processing logs will be visible in the program’s terminal output.
3.	Output Topic: The processed result will be published to the output topic. Check the terminal subscribed to the output topic for the result.

---

### Troubleshooting

**1.	MQTT Connection Errors:** 
verify the broker URL and port in your .env file or command line arguments match.

**2.  Invalid Payload:** ensure the payload adheres to the expected JSON structure

```json
{
  "id": "test-id",
  "numberOfChildren": 2,
  "familyComposition": "couple",
  "familyUnitInPayForDecember": true
}
```
**Notes:** the program assumes `numberOfChildren >= 0` and validates the `familyComposition` field with either values of `"single"` or `"couple"` along with requiring `id` field to not be empty and `familyUnitsPayForDecember` to be boolean. 

Any invalid input message will be displayed in the terminal listening to the input topic, however it will not be processed and the output will not be published to the output topic

**3. Web app issues:** if the original web app publisher does not seem to send payload to the input topic id, confirm its connection to the MQTT broker

**4. Internet Connection:** Ensure connection to the internet is strong and secure

---

## Running Unit Tests
The project includes unit tests to ensure the correctness of the rule engine and other key components. Follow these steps to run the tests:
1. **Run all tests:** in the root folder, use the following command to execute all unit tets in the project:
```git
go test ./... -v
```
-  ./... ensures that test in all subdirectories are run
- -v provides verbose output to display detailed information about the test results
---
2. **Run tests for a specific package:** if you want to run test for a specific package (e.g. the rules package), use:
```git
go test ./rules -v
```

---

3. **Run a specific test function:** to run a specific test function (e.g. TestCalculateSupplement), use the --run flag:
```git
go test ./rules -v -run TestCalculateSupplement
```
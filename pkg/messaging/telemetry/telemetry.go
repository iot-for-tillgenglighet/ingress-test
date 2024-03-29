package telemetry

import "github.com/iot-for-tillgenglighet/ingress-test/pkg/messaging"

// Temperature is a telemetry type IoTHubMessage
type Temperature struct {
	messaging.IoTHubMessage
	Temp float64 `json:"temp"`
}

// ContentType returns the ContentType for a Temperature telemetry message
func (msg *Temperature) ContentType() string {
	return "application/json"
}

// TopicName returns the name of the topic that a Temperature telemetry message should be posted to
func (msg *Temperature) TopicName() string {
	return "telemetry.temperature"
}

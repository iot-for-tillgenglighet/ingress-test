package main

import (
	"github.com/iot-for-tillgenglighet/ingress-test/pkg/handler"
	"github.com/iot-for-tillgenglighet/ingress-test/pkg/messaging"

	"os"
	"time"

	log "github.com/sirupsen/logrus"
	
)

func main() {

	serviceName := "ingress-test"

	rabbitMQHostEnvVar := "RABBITMQ_HOST"
	rabbitMQHost := os.Getenv(rabbitMQHostEnvVar)
	rabbitMQUser := os.Getenv("RABBITMQ_USER")
	rabbitMQPass := os.Getenv("RABBITMQ_PASS")

	if rabbitMQHost == "" {
		log.Fatal("Rabbit MQ host missing. Please set " + rabbitMQHostEnvVar + " to a valid host name or IP.")
	}

	var messenger *messaging.Context
	var err error

	for messenger == nil {

		time.Sleep(2 * time.Second)

		messenger, err = messaging.Initialize(messaging.Config{
			ServiceName: serviceName,
			Host:        rabbitMQHost,
			User:        rabbitMQUser,
			Password:    rabbitMQPass,
		})

		if err != nil {
			log.Error(err)
		}
	}

	defer messenger.Close()
	
	
	handler.InitializeRouter(messenger)

}

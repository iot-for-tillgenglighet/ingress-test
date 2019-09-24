package handler

import (
	"github.com/iot-for-tillgenglighet/ingress-test/pkg/messaging"
	"github.com/iot-for-tillgenglighet/ingress-test/pkg/messaging/telemetry"

	"log"
	"net/http"
	"os"
	"bytes"
	"strings"
	"encoding/json"


	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var messenger *messaging.Context

func InitializeRouter(m *messaging.Context) {

	messenger = m
	router := mux.NewRouter()


	router.HandleFunc("/", getMethod).Methods("GET")
	router.HandleFunc("/put", putMethod).Methods("PUT")

	port := os.Getenv("TEST_API_PORT")
	if port == "" {
		port = "8888"
	}

	log.Printf("Starting ingress-test on port %s, \n", port)

	err := http.ListenAndServe(":"+port, handlers.LoggingHandler(os.Stdout, router))
	if err != nil {
		log.Print(err)
	}

	
}

func getMethod(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello"))

}

func putMethod(w http.ResponseWriter, r *http.Request) {

	requestBuffer := new(bytes.Buffer)
	requestBuffer.ReadFrom(r.Body)
	
	requestBody := requestBuffer.String()
	lines := strings.Split(requestBody, "\n")

	for _, l := range lines {
		parts := strings.SplitN(l, "\t", 2)
		topicName := parts[0]
		messageData := parts[1]

		message := &telemetry.Temperature{}

		if topicName == "telemetry.temperature" {
			_ = json.Unmarshal([]byte(messageData), message)
			messenger.PublishOnTopic(message)
		}
	
	}
	
}

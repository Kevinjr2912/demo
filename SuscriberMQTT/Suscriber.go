package suscribermqtt

import (
	"fmt"
	"log"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/subosito/gotenv"
)

type Suscriber struct {
	client   mqtt.Client
	topic    string
}

func NewSuscriberMQTT() *Suscriber {

	// Cargar variables de entorno
	err := gotenv.Load()
	if err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
	}

	brokerURL := os.Getenv("MQTTBROKER")
	topic := os.Getenv("TOPIC")

	client := CreateClient(brokerURL)

	suscriber := &Suscriber{
		client: client,
		topic:  topic,
	}

	suscriber.Subscribe()

	return suscriber

}

func (s *Suscriber) Subscribe() {

	if token := s.client.Subscribe(s.topic, 0, messageHandler); token.Wait() && token.Error() != nil {
		log.Fatalf("Error al suscribirse al tópico: %v", token.Error())
	}

	fmt.Println("Suscrito al tópico MQTT. Esperando mensajes...")

}

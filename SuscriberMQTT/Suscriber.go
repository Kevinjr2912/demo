package suscribermqtt

import (
	"fmt"
	"log"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gorilla/websocket"
	"github.com/subosito/gotenv"
)

type Suscriber struct {
	client         mqtt.Client
	topic          string
	messageHandler mqtt.MessageHandler
	wsConnection   *websocket.Conn
}

func NewSuscriberMQTT(wsConnection *websocket.Conn) *Suscriber {

	// Cargar variables de entorno
	err := gotenv.Load()
	if err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
	}

	brokerURL := os.Getenv("MQTTBROKER")
	topic := os.Getenv("TOPIC")

	suscriber := &Suscriber{}

	var messageHandler mqtt.MessageHandler = suscriber.HandleMessage

	client := CreateClient(brokerURL, suscriber.messageHandler)

	suscriber.client = client
	suscriber.topic = topic
	suscriber.messageHandler = messageHandler
	suscriber.wsConnection = wsConnection

	suscriber.Subscribe()

	return suscriber

}

func (s *Suscriber) Subscribe() {
	if token := s.client.Subscribe(s.topic, 0, s.messageHandler); token.Wait() && token.Error() != nil {
		log.Fatalf("Error al suscribirse al tópico: %v", token.Error())
	}

	fmt.Println("Suscrito al tópico MQTT. Esperando mensajes...")

}

func (s *Suscriber) HandleMessage(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Mensaje recibido en %s: %s\n", msg.Topic(), msg.Payload())

	processSensorData(msg.Payload(), s.wsConnection)
}

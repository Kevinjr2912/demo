package suscribermqtt

import (
	"log"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// CreateClient crea y devuelve un cliente MQTT
func CreateClient(brokerURL string) mqtt.Client {
	
	opts := mqtt.NewClientOptions()
	opts.AddBroker(brokerURL)
	opts.SetClientID("go_mqtt_subscriber")
	opts.SetDefaultPublishHandler(messageHandler)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Error al conectar con MQTT: %v", token.Error())
	}

	return client

}

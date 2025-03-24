// package main

// import (
// 	"fmt"
// 	"log"

// 	mqtt "github.com/eclipse/paho.mqtt.golang"
// )

// // Configuración del broker MQTT
// const (
// 	mqttBroker = "tcp://18.214.187.219:1883"
// 	mqttTopic  = "test/topic"
// )

// // Función que maneja los mensajes recibidos
// var messageHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
// 	fmt.Printf("Mensaje recibido en %s: %s\n", msg.Topic(), msg.Payload())
// }

// func main() {
// 	// Configurar cliente MQTT
// 	opts := mqtt.NewClientOptions()
// 	opts.AddBroker(mqttBroker)
// 	opts.SetClientID("go_mqtt_subscriber")
// 	opts.SetDefaultPublishHandler(messageHandler)

// 	// Crear cliente MQTT
// 	client := mqtt.NewClient(opts)
// 	if token := client.Connect(); token.Wait() && token.Error() != nil {
// 		log.Fatalf("Error al conectar con MQTT: %v", token.Error())
// 	}

// 	// Suscribirse al tópico
// 	if token := client.Subscribe(mqttTopic, 0, nil); token.Wait() && token.Error() != nil {
// 		log.Fatalf("Error al suscribirse al tópico: %v", token.Error())
// 	}

// 	fmt.Println("Suscrito al tópico MQTT. Esperando mensajes...")

// 	// Mantener la ejecución para seguir recibiendo mensajes
// 	select {}
// }
package main

import (
	"log"
	"net/http"
	"test/SuscriberMQTT"
)

func main() {

	// Inicializa el suscriptor MQTT
	suscribermqtt.NewSuscriberMQTT()

	// Registramos las rutas
	// routes.RegisterRoutes()

	// Levantamos el servidor HTTP en un goroutine
	go func() {
		log.Println("Iniciando servidor HTTP en el puerto 8080...")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatalf("Error al iniciar el servidor: %v", err)
		}
	}()

	
	select {}
}

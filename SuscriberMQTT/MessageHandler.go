package suscribermqtt

import (
	"encoding/json"
	"fmt"
	"test/SuscriberMQTT/validators"
	"test/models"
	"test/request"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var dataIoT = &models.DataIoT{}

var messageHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Mensaje recibido en %s: %s\n", msg.Topic(), msg.Payload())

	// Procesar los datos de los sensores
	processSensorData(msg.Payload())
}

func processSensorData(data []byte) {
	var dataIoTLocal models.DataIoT

	if err := json.Unmarshal(data, &dataIoTLocal); err != nil {
		fmt.Errorf("Error al deserializar el mensaje: %v", err)
	}

	validators.ValidateData(&dataIoTLocal, dataIoT)

	verifyValues()
}

func verifyValues() {

	if dataIoT.IdPlot != 0 && (dataIoT.Temperature != 0 || dataIoT.Temperature == 0) && dataIoT.AirQuality != 0 && (dataIoT.Humidity != 0 || dataIoT.Humidity == 0) {
		
		fmt.Printf("Id de la parcela: %v\n", dataIoT.IdPlot)
		fmt.Printf("Temperatura: %v\n", dataIoT.Temperature)
		fmt.Printf("Calidad del aire: %v\n", dataIoT.AirQuality)
		fmt.Printf("Humedad: %v\n", dataIoT.Humidity)

		// Pendiente
		dataIoT.Sun = 50.02

		data, err := json.Marshal(dataIoT)

		if err != nil {
			fmt.Errorf("Error", err.Error())
		}


		// Hacer la petición a la API
		request.Fetch(data)

		// Enviar al ws
		
	}

}

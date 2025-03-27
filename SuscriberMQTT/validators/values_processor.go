package validators

import (
	"fmt"
	"test/models"
)

func validateTemperature(temp float32) bool {
	return temp >= 0 && temp <= 50
}

func validateHumidity(humidity float32) float32 {
	if humidity < 0 {
		return 0
	} else if humidity > 100 {
		return 100
	}
	return humidity
}

func ValidateData(dataIoTLocal, dataIoT *models.DataIoT) {
	
	// Verificar id
	if dataIoTLocal.IdPlot > 0 && dataIoTLocal.IdPlot != dataIoT.IdPlot {
		fmt.Println("Exit id")
		dataIoT.IdPlot = dataIoTLocal.IdPlot
	}

	// Verificar temperatura
	if validateTemperature(dataIoTLocal.Temperature) && dataIoT.Temperature != 0 {
		fmt.Println("Exit temp")
		dataIoT.Temperature = dataIoTLocal.Temperature
	}

	// Verificar humedad
	dataIoTLocal.Humidity = validateHumidity(dataIoTLocal.Humidity)
	
	if dataIoT.Humidity == 0 {
		fmt.Println("Exit humedad")
		dataIoT.Humidity = dataIoTLocal.Humidity
	}

	// Verificar la luz solar
	dataIoT.Sun = dataIoTLocal.Sun

	// Verificar la calidad del aire
	dataIoT.AirQuality = dataIoTLocal.AirQuality

	fmt.Println("Antes de enviarse", dataIoT)
	

}
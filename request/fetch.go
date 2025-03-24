package request

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

func Fetch(data []byte) {

	url := "http://localhost:8080/measurements/"

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))

	if err != nil {
		fmt.Errorf("Error al hacer la petición")
	}

	if resp.StatusCode == http.StatusCreated {
		log.Println("Datos enviados a la API")
	}

}
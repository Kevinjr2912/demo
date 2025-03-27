package main

import (
	"fmt"
	"net/http"
	"test/controllers"
)

func main() {

	http.HandleFunc("/ws", controllers.Echo)

    fmt.Println("WebSocket server started on :8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println("ListenAndServe:", err)
    }

}

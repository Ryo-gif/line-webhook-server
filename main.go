package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	log.Println("Webhook received:", string(body))
	fmt.Fprintln(w, "OK")
}

func main() {
	http.HandleFunc("/webhook", webhookHandler)
	log.Println("Server started on :10000")
	log.Fatal(http.ListenAndServe(":10000", nil))
}

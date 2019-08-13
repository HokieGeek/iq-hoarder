package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/hokiegeek/iqhoarder"
)

func main() {
	portPtr := flag.Int("port", 8078, "Port to run the daemon on")
	// secretPtr := flag.String("secret", "secret123!", "The secrit")

	flag.Parse()

	http.HandleFunc("/hoard", iqhoarder.IQWebhookHandler)

	log.Printf("Starting iq-hoarder on port %d\n", *portPtr)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *portPtr), nil))
}

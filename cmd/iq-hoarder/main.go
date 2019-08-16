package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	// "github.com/hokiegeek/iqhoarder"
	iqhooks "github.com/sonatype-nexus-community/gonexus/iq/webhooks"
)

func main() {
	portPtr := flag.Int("port", 8078, "Port to run the daemon on")
	// secretPtr := flag.String("secret", "secret123!", "The secrit")

	flag.Parse()

	// appEvalEvents, _ := iqhooks.ApplicationEvaluationEvents()
	violationAlertEvents, _ := iqhooks.ViolationAlertEvents()

	go func() {
		for {
			select {
			// case <-appEvalEvents:
			// 	log.Println("<APP EVAL>")
			case violation := <-violationAlertEvents:
				log.Println("<VIOLATION> ReportID: ", violation.ApplicationEvaluation.ReportID)
			default:
			}
		}
	}()

	http.HandleFunc("/hoard", iqhooks.Listen)

	log.Printf("Starting iq-hoarder on port %d\n", *portPtr)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *portPtr), nil))
}

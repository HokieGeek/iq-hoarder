package main

import (
	"flag"
	"time"

	"github.com/hokiegeek/iqhoarder"
	"github.com/hokiegeek/iqhoarder/couchdb"
)

func main() {
	// portPtr := flag.Int("port", 8078, "Port to run the daemon on")
	// iqServerPtr := flag.String("iq", "http://localhost:8070", "")
	// iqAuthPtr := flag.String("iqAuth", "admin:admin123", "")
	// connectorPtr := flag.String("connector", "couchdb", "Specify the connector to use as a database")
	// secretPtr := flag.String("secret", "secret123!", "The secrit")

	flag.Parse()

	db, err := couchdb.New("hoarder", "foobar")
	if err != nil {
		panic(err)
	}

	db.Insert(iqhoarder.Report{ID: "dummyID", ApplicationPublicID: "appID", Created: time.Now()})

	/*
		// appEvalEvents, _ := iqhooks.ApplicationEvaluationEvents()
		violationAlertEvents, _ := iqhooks.ViolationAlertEvents()

		go func() {
			for {
				select {
				// case <-appEvalEvents:
				// 	log.Println("<APP EVAL>")
				case violation := <-violationAlertEvents:
					log.Println("<VIOLATION> ReportID: ", violation.ApplicationEvaluation.ReportID)
					// iqhoarder.Add(db, foo)
				default:
				}
			}
		}()

		http.HandleFunc("/hoard", iqhooks.Listen)
		http.HandleFunc("/reports", TODO)

		log.Printf("Starting iq-hoarder on port %d\n", *portPtr)
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *portPtr), nil))
	*/
}

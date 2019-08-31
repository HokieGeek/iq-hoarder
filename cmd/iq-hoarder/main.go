package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	// "github.com/sonatype-nexus-community/gonexus/iq/iqwebhooks"
	"github.com/hokiegeek/iqhoarder"
	"github.com/hokiegeek/iqhoarder/couchdb"
)

func main() {
	portPtr := flag.Int("port", 8078, "Port to run the daemon on")
	// iqServerPtr := flag.String("iq", "http://localhost:8070", "")
	// iqAuthPtr := flag.String("iqAuth", "admin:admin123", "")
	// connectorPtr := flag.String("connector", "couchdb", "Specify the connector to use as a database")
	// secretPtr := flag.String("secret", "secret123!", "The secrit")

	flag.Parse()

	var (
		db  iqhoarder.DB
		err error
	)

	db, err = couchdb.New("hoarder", "http://localhost:8790")
	if err != nil {
		panic(err)
	}

	// db.Insert(iqhoarder.Report{ID: "dummyID", Server: *iqServerPtr, ApplicationPublicID: "appID", CreationTime: time.Now()})

	/*
		appEvalEvents, _ := iqwebhooks.ApplicationEvaluationEvents()
		violationAlertEvents, _ := iqwebhooks.ViolationAlertEvents()

		go func() {
			for {
				select {
				case appEval := <-appEvalEvents:
					log.Println("<APP EVAL>")
					buf, err := json.MarshalIndent(appEval, "", "  ")
					if err != nil {
						log.Println("foo")
						continue
					}
					log.Println(string(buf))
					// db.Insert(db, foo)
				// case violation := <-violationAlertEvents:
					// log.Println("<VIOLATION> ReportID: ", violation.ApplicationEvaluation.ReportID)
					// db.Insert(db, foo)
				default:
				}
			}
		}()

		http.HandleFunc("/api/v1/hoard", iqwebhooks.Listen)

		http.HandleFunc("/api/v1/reports", func(w http.ResponseWriter, r *http.Request) {
			iqhoarder.APIReports(w, r, db)
		})
	*/

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		iqhoarder.HTMLReportsList(w, r, db)
	})

	log.Printf("Starting iq-hoarder on port %d\n", *portPtr)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *portPtr), nil))
}

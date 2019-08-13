package iqhoarder

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/sonatype-nexus-community/gonexus/iq"
)

// IQWebhookHandler handles an incoming webhook from IQ Server
func IQWebhookHandler(w http.ResponseWriter, r *http.Request) {
	ok, whtype := nexusiq.IsWebhookEvent(r)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch whtype {
	case nexusiq.WebhookEventApplicationEvaluation:
		log.Println("Accepted Application Evaluation")

		var wh nexusiq.WebhookApplicationEvaluation
		err = json.Unmarshal(body, &wh)

		fmt.Println(wh)
	case nexusiq.WebhookEventPolicyAlert:
		log.Println("Accepted Policy Alert")

		var wh nexusiq.WebhookViolationAlert
		err = json.Unmarshal(body, &wh)

		fmt.Println(wh)
	default:
		log.Println(whtype)
		log.Println(string(body))
	}
}

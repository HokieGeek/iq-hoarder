package iqhoarder

import (
	"fmt"
	"html"
	"net/http"
)

// APIReports provides a REST api for interacting with the reports in the database
func APIReports(w http.ResponseWriter, r *http.Request, db DB) {
	fmt.Fprintf(w, "API TODO: %q", html.EscapeString(r.URL.Path))
}

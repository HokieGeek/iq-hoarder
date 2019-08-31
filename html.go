package iqhoarder

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"time"
)

var pageTmpl = `<html>
	<head>
		<title>{{ .Title }}</title>
	</head>
	<body>
		<ul>
		{{ range $key, $value := .ReportsByApp }}
			<li><a href="http://todo">{{ $key }}: {{ $value.ID }}</a></li>
		{{ end }}
		</ul>
	</body>
</html>`

// WriteHTML writes the reports in the database to an io.Writer after applying an HTML template
func WriteHTML(w io.Writer, db DB) error {
	_, err := db.Query(NewQueryBuilder())
	if err != nil {
		return err
	}

	page := struct {
		Title        string
		ReportsByApp map[string]Report
	}{
		Title: "foo",
		ReportsByApp: map[string]Report{
			"app1": {ID: "dummyID1", Server: "http://iq:1111", ApplicationPublicID: "app1", CreationTime: time.Now()},
			"app2": {ID: "dummyID2", Server: "http://iq:2222", ApplicationPublicID: "app2", CreationTime: time.Now()},
		},
	}

	t := template.Must(template.New("page").Parse(pageTmpl))
	if err := t.Execute(w, page); err != nil {
		return fmt.Errorf("could not generate page: %v", err)
	}

	return nil
}

// HTMLReportsList serves a page with all of the reports in the database
func HTMLReportsList(w http.ResponseWriter, r *http.Request, db DB) {
	if err := WriteHTML(w, db); err != nil {
		log.Printf("could not serve page: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

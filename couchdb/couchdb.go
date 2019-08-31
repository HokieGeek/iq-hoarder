package couchdb

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/hokiegeek/iqhoarder"
)

type reportDoc struct {
	ID                  string    `json:"_id"`
	Server              string    `json:"server"`
	ApplicationPublicID string    `json:"applicationPublidId"`
	CreationTime        time.Time `json:"creationTime"`
}

func makeDoc(report iqhoarder.Report) reportDoc {
	return reportDoc{
		ID:                  report.ID,
		Server:              report.Server,
		ApplicationPublicID: report.ApplicationPublicID,
		CreationTime:        report.CreationTime,
	}
}

func fromDoc(doc reportDoc) iqhoarder.Report {
	return iqhoarder.Report{
		ID:                  doc.ID,
		Server:              doc.Server,
		ApplicationPublicID: doc.ApplicationPublicID,
		CreationTime:        doc.CreationTime,
	}
}

type couchDB struct {
	databaseName, host string
}

func (db *couchDB) req(method, endpoint string, data io.Reader, headers map[string]string) (*http.Response, error) {
	url := fmt.Sprintf("%s/%s/%s", db.host, db.databaseName, endpoint)
	req, err := http.NewRequest(method, url, data)
	if err != nil {
		return nil, fmt.Errorf("could not create request: %v", err)
	}

	// req.Header.Set("Content-Length", fmt.Sprintf("%d", data.Len()))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	return client.Do(req)
}

func (db *couchDB) reportPath(r iqhoarder.Report) string {
	// return fmt.Sprintf("%s/%s", db.databaseName, r.ID)
	return r.ID
}

func (db *couchDB) revisionID(r iqhoarder.Report) (revID string, err error) {
	resp, err := db.req(http.MethodHead, db.reportPath(r), nil, nil)
	if err != nil {
		return
	}

	switch resp.StatusCode {
	case http.StatusOK, http.StatusNotModified:
		revID = resp.Header.Get("Etag")
	default:
		err = errors.New("revision ID not found")
	}

	return
}

// Insert adds a new report to the CouchDB database
func (db *couchDB) Insert(report iqhoarder.Report) error {
	buf, err := json.Marshal(makeDoc(report))
	if err != nil {
		return err // TODO
	}

	// resp, err := db.req(http.MethodPost, db.databaseName, bytes.NewBuffer(buf), nil)
	resp, err := db.req(http.MethodPost, "", bytes.NewBuffer(buf), nil)
	if err != nil {
		return err // TODO
	}

	if resp.StatusCode == http.StatusConflict {
		return errors.New("TODO") // TODO
	}

	// fmt.Println("Insert")
	return nil
}

// Query returns reports which match the search criteria
func (db *couchDB) Query(query *iqhoarder.QueryBuilder) ([]iqhoarder.Report, error) {
	var (
		resp *http.Response
		err  error
		r    io.Reader
	)

	selectors := query.Build()
	// TODO: query
	if len(selectors) == 0 {
		r = strings.NewReader(`{ "selector": { "_id": { "$gt": null } } }`)
	}
	resp, err = db.req(http.MethodPost, "_find", r, nil)
	if err != nil {
		return nil, err // TODO
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err // TODO
	}

	fmt.Println("Query: ", string(body))

	return nil, nil
}

// New creates a new instance creates a new instance of CouchDB
func New(databaseName, host string) (iqhoarder.DB, error) {
	db := couchDB{databaseName, host}

	// check if DB exists
	resp, err := db.req(http.MethodHead, "", nil, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to check if database exists: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		// Create if it is not found
		resp, err = db.req(http.MethodPut, "", nil, nil)
		if err != nil {
			return nil, fmt.Errorf("error creating database: %v", err)
		}
		if resp.StatusCode != http.StatusCreated {
			return nil, fmt.Errorf("did not create database: %s", resp.Status)
		}
	}

	return &db, nil
}

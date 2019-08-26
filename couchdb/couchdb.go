package couchdb

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/hokiegeek/iqhoarder"
)

type couchDB struct {
	databaseName, host string
}

func (db *couchDB) req(method, endpoint string, data io.Reader, headers map[string]string) (*http.Response, error) {
	url := fmt.Sprintf("%s/%s", db.host, endpoint)
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
	return fmt.Sprintf("%s/%s", db.databaseName, r.ID)
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
	return nil
}

// Query returns reports which match the search criteria
func (db *couchDB) Query(query iqhoarder.QueryBuilder) ([]iqhoarder.Report, error) {
	// TODO
	return nil, nil
}

// New creates a new instance creates a new instance of CouchDB
func New(databaseName, host string) (iqhoarder.DB, error) {
	db := couchDB{databaseName, host}

	// check if DB exists
	resp, err := db.req(http.MethodHead, db.databaseName, nil, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		// Create if it is not found
		resp, err = db.req(http.MethodPut, db.databaseName, nil, nil)
		if err != nil {
			return nil, err
		}
		if resp.StatusCode != http.StatusOK {
			return nil, errors.New("foobar")
		}
	}

	return &db, nil
}

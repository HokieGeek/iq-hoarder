package iqhoarder

import (
	"time"
)

// Report defines the information which defines an IQ report
type Report struct {
	Server, ID, ApplicationPublicID string
	CreationTime                    time.Time
}

// QueryBuilder TODO
type QueryBuilder struct {
	selectors map[string]string
}

// Stage adds a pipeline stage to query by
func (b *QueryBuilder) Stage(s string) *QueryBuilder {
	b.selectors["stage"] = s
	return b
}

// Application adds an application name to query by
func (b *QueryBuilder) Application(n string) *QueryBuilder {
	b.selectors["applicationPublicID"] = n
	return b
}

// Build returns the selectors map
func (b *QueryBuilder) Build() map[string]string {
	return b.selectors
}

// NewQueryBuilder creates a new instance of QueryBuilder
func NewQueryBuilder() *QueryBuilder {
	b := new(QueryBuilder)
	b.selectors = make(map[string]string)
	return b
}

// DB defines the interface which a database client needs to comply with
type DB interface {
	Insert(report Report) error
	Query(query *QueryBuilder) ([]Report, error)
}

package model

import (
	"time"
)

type TraceRecord struct {
	Ip        string //`json:"ip"`
	Uri       string //`json:"uri"`
	Datetime  string
	Refer     string
	UserAgent string
	Time      time.Time //`json:"time"`
}

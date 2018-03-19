package model

import (
	"time"
)

type TraceRecord struct {
	Ip       string //`json:"ip"`
	Uri      string //`json:"uri"`
	Datetime string
	Time     time.Time //`json:"time"`
}

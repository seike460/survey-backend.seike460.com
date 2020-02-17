package models

import (
	"time"

	"github.com/rs/xid"
)

// Survey  SurveyData
type Survey struct {
	UUID xid.ID    // Hash key
	Time time.Time // Range key
	Msg  string    `dynamo:"Message"`
}

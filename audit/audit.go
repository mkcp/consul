// TODO Read VLT-082 https://docs.google.com/document/d/1fGVhNX_Y6lSeCUvRUq22y7Q85u6kmMzK94IF-Ws-nsE/edit
// TODO Read [PRD-DRAFT CSL-007] https://docs.google.com/document/d/1f_hGqDDtb1aij5R3I5KiAgYYfDSyJnEViuqYUv1tM5E/edit
// TODO Stub out CSL-XXX Audit Logging MVP https://docs.google.com/document/d/1KqYsHsRBekkZ1ksdX4oAP_kHz9ihmoSpfzi6nOgF3Ko/edit#heading=h.pkdh1kq4nz8z
// TODO Read Nomad's WIP auditor https://github.com/hashicorp/nomad-enterprise/compare/wip/audit-logging
package audit

import (
	"fmt"
	"github.com/hashicorp/eventlogger"
	"time"
)

// Auditor is lifted from Nomad's structure
type Auditor struct {
	broker eventlogger.Broker
	et     eventlogger.EventType
}

// Audit is a WIP
func NewAuditor(eventType string) Auditor {
	return Auditor{
		broker: *eventlogger.NewBroker(),
		et:     eventlogger.EventType(eventType),
	}
}

// TODO Other kinds of sinks for the eventlogger?
type Config struct {
	LogFile string
}

// Maybe this is area specific fields? Can we define this at a top level?
type Event struct {
	ID          string
	Stage       string
	Operation   string
	Path        string
	DisplayName string
	Timestamp   string
	Request     interface{}
	Response    interface{}
}

type Eventer interface {
	Emit(eventType string, event Event)
}

func NewEvent(input interface{}) eventlogger.Event {
	et := eventlogger.EventType(fmt.Sprint(input))
	event := eventlogger.Event{
		Type:      et,
		CreatedAt: time.Now(),
		Formatted: nil,
		Payload:   nil,
	}
	return event
}

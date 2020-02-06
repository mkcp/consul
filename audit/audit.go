// TODO Read VLT-082 https://docs.google.com/document/d/1fGVhNX_Y6lSeCUvRUq22y7Q85u6kmMzK94IF-Ws-nsE/edit
// TODO Read [PRD-DRAFT CSL-007] https://docs.google.com/document/d/1f_hGqDDtb1aij5R3I5KiAgYYfDSyJnEViuqYUv1tM5E/edit
// TODO Stub out CSL-XXX Audit Logging MVP https://docs.google.com/document/d/1KqYsHsRBekkZ1ksdX4oAP_kHz9ihmoSpfzi6nOgF3Ko/edit#heading=h.pkdh1kq4nz8z
// TODO Read Nomad's WIP auditor https://github.com/hashicorp/nomad-enterprise/compare/wip/audit-logging
package audit

import (
	"github.com/hashicorp/eventlogger"
	"time"
)

// Auditor is lifted from Nomad's structure
type Auditor struct {
	broker eventlogger.Broker
	et eventlogger.EventType
}

// Audit is a WIP
func Audit(input interface{}) interface{} {
	event := eventlogger.Event{
		Type: "TestEvent",
		CreatedAt: time.Now(),
		Formatted: nil,
		Payload:   nil,
	}
	return event
}

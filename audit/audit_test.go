package audit

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/eventlogger"
	"testing"
)

func TestAudit_Broker(t *testing.T) {
}

type testCase struct {
	name    string
	auditor Auditor
}

func testCases() []testCase {
	return []testCase{
		{
			"first",
			NewAuditor("eventFirst"),
		},
		{
			"second",
			NewAuditor("eventSecond"),
		},
	}
}

func TestAudit(t *testing.T) {
	t.Run("Audit test", func(t *testing.T) {
		event := Audit("TestEvent")
		e := spew.Sdump(event)
		fmt.Printf("[TESTING] audit event=%v", e)
	})
}

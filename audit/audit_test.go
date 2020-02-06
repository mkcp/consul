package audit

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func TestAudit(t *testing.T) {
	t.Run("Audit test", func(t *testing.T){
		event := Audit(nil)
		e := spew.Sdump(event)
		fmt.Printf("[TESTING] audit event=%v", e)
	})
}

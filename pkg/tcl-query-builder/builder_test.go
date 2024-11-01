package tcl_query_builder

import (
	"testing"
)

func TestBuilder(t *testing.T) {
	qb := NewQueryBuilder()
	qb.AddCondition(Eq("WorkflowType", "order_processing"))
	qb.AddCondition(Eq("Status", 123))
	qb.AddCondition(Gte("StartTime", "2021-01-01"))
	qb.AddCondition(Lte("StartTime", "2021-01-02"))

	qb.Order("StartTime", true)
	qb.Order("DisplayID", false)

	want := "WorkflowType = 'order_processing' and Status = 123 and StartTime >= '2021-01-01' and StartTime <= '2021-01-02' order by StartTime asc, DisplayID desc"
	if got := qb.Query(); got != want {
		t.Errorf("Query() = %v, want %v", got, want)
	}
}

package pkg

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func CheckIsNumber(t *testing.T, s string, m ...interface{}) {
	n, err := strconv.Atoi(s)
	assert.Nil(t, err, "Expected string to be a number")
	s2 := strconv.Itoa(n)
	assert.Equal(t, s, s2, "Expected stringified number to equal the original string")
}

func CheckNotBlank(t *testing.T, s string, m ...interface{}) {
	v := strings.TrimSpace(s)
	assert.NotEmpty(t, v, m)
}

func CheckWorkItem(t *testing.T, w WorkItem) {
	CheckNotBlank(t, w.Title, "WorkItem.Title")
	CheckNotBlank(t, w.Description, "WorkItem.Description")
	CheckNotBlank(t, w.Work_item_id, "WorkItem.Work_item_id")
	CheckIsNumber(t, w.Work_item_id, "WorkItem.Work_item_id")
	CheckNotBlank(t, w.Tag_id, "WorkItem.Tag_id")
	CheckNotBlank(t, w.Status_id, "WorkItem.Status_id")
}

func CheckOrder(t *testing.T, o WorkItem) {
	CheckWorkItem(t, o)
}

func CheckBatch(t *testing.T, b WorkItem) {
	CheckWorkItem(t, b)
	CheckNotBlank(t, b.Parent_work_item_id, "WorkItem.Parent_work_item_id")
	CheckIsNumber(t, b.Parent_work_item_id, "WorkItem.Parent_work_item_id")
}

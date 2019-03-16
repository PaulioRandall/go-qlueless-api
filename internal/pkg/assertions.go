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
	CheckNotBlank(t, w.WorkItemID, "WorkItem.Work_item_id")
	CheckIsNumber(t, w.WorkItemID, "WorkItem.Work_item_id")
	CheckNotBlank(t, w.TagID, "WorkItem.Tag_id")
	CheckNotBlank(t, w.StatusID, "WorkItem.Status_id")
}

func CheckOrder(t *testing.T, o WorkItem) {
	CheckWorkItem(t, o)
}

func CheckBatch(t *testing.T, b WorkItem) {
	CheckWorkItem(t, b)
	CheckNotBlank(t, b.ParentWorkItemID, "WorkItem.Parent_work_item_id")
	CheckIsNumber(t, b.ParentWorkItemID, "WorkItem.Parent_work_item_id")
}

func CheckTag(t *testing.T, e map[string]interface{}) {
	assert.NotNil(t, e["title"])
	CheckNotBlank(t, e["title"].(string), "Dicts.Tag.Title")
	assert.NotNil(t, e["description"])
	CheckNotBlank(t, e["description"].(string), "Dicts.Tag.Description")
	assert.NotNil(t, e["tag_id"])
	CheckNotBlank(t, e["tag_id"].(string), "Dicts.Tag.Tag_id")
}

func CheckStatus(t *testing.T, e map[string]interface{}) {
	assert.NotNil(t, e["title"])
	CheckNotBlank(t, e["title"].(string), "Dicts.Status.Title")
	assert.NotNil(t, e["description"])
	CheckNotBlank(t, e["description"].(string), "Dicts.Tag.Description")
	assert.NotNil(t, e["status_id"])
	CheckNotBlank(t, e["status_id"].(string), "Dicts.Tag.Status_id")
}

func CheckWorkItemType(t *testing.T, e map[string]interface{}) {
	assert.NotNil(t, e["title"])
	CheckNotBlank(t, e["title"].(string), "Dicts.WorkItemType.Title")
	assert.NotNil(t, e["description"])
	CheckNotBlank(t, e["description"].(string), "Dicts.WorkItemType.Description")
	assert.NotNil(t, e["work_item_type_id"])
	CheckNotBlank(t, e["work_item_type_id"].(string), "Dicts.WorkItemType.Work_item_type_id")
}

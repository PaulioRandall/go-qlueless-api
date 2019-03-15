package dictionaries

import (
	"testing"

	shr "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
	"github.com/stretchr/testify/assert"
)

// When invoked, should not return nil
func TestLoadDictReply___1(t *testing.T) {
	assert.NotNil(t, LoadDictsReply())
}

// When invoked, should return valid reply struct containing valid dictionaries
// with entries
func TestLoadDictReply___2(t *testing.T) {
	act := LoadDictsReply()

	tags := act.Data.(map[string]interface{})["tags"].([]interface{})
	assert.NotNil(t, tags)
	for _, e := range tags {
		shr.CheckTag(t, e.(map[string]interface{}))
	}

	statuses := act.Data.(map[string]interface{})["statuses"].([]interface{})
	assert.NotNil(t, statuses)
	for _, e := range statuses {
		shr.CheckStatus(t, e.(map[string]interface{}))
	}

	workItemTypes := act.Data.(map[string]interface{})["work_item_types"].([]interface{})
	assert.NotNil(t, workItemTypes)
	for _, e := range workItemTypes {
		shr.CheckWorkItemType(t, e.(map[string]interface{}))
	}
}

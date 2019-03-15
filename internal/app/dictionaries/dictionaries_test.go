package dictionaries

import (
	"testing"

	shr "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
	"github.com/stretchr/testify/assert"
)

// When invoked, should not return nil
func TestLoadDict___1(t *testing.T) {
	assert.NotNil(t, LoadDicts())
}

// When invoked, should return valid reply struct containing valid dictionaries
// with entries
func TestLoadDicts___2(t *testing.T) {
	act := LoadDicts()

	tags := act["tags"].([]interface{})
	assert.NotNil(t, tags)
	for _, e := range tags {
		shr.CheckTag(t, e.(map[string]interface{}))
	}

	statuses := act["statuses"].([]interface{})
	assert.NotNil(t, statuses)
	for _, e := range statuses {
		shr.CheckStatus(t, e.(map[string]interface{}))
	}

	workItemTypes := act["work_item_types"].([]interface{})
	assert.NotNil(t, workItemTypes)
	for _, e := range workItemTypes {
		shr.CheckWorkItemType(t, e.(map[string]interface{}))
	}
}

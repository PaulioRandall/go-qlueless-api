package dictionaries

import (
	"testing"

	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
	"github.com/stretchr/testify/assert"
)

// When invoked, should not return nil
func TestLoadDict___1(t *testing.T) {
	LoadDicts()
	assert.NotNil(t, dicts)
}

// When invoked, should create a dictionary map containing a valid tag
// dictionary
func TestLoadDicts___2(t *testing.T) {
	LoadDicts()

	tags := dicts["tags"].([]interface{})
	assert.NotNil(t, tags)
	for _, e := range tags {
		CheckTag(t, e.(map[string]interface{}))
	}
}

// When invoked, should create a dictionary map containing a valid statuses
// dictionary
func TestLoadDicts___3(t *testing.T) {
	LoadDicts()

	statuses := dicts["statuses"].([]interface{})
	assert.NotNil(t, statuses)
	for _, e := range statuses {
		CheckStatus(t, e.(map[string]interface{}))
	}
}

// When invoked, should create a dictionary map containing a valid thing
// types dictionary
func TestLoadDicts___4(t *testing.T) {
	LoadDicts()

	workItemTypes := dicts["work_item_types"].([]interface{})
	assert.NotNil(t, workItemTypes)
	for _, e := range workItemTypes {
		CheckThingType(t, e.(map[string]interface{}))
	}
}

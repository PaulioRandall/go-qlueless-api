//usr/bin/env go run "$0" "$@"; exit "$?"

package api

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type Thing struct {
	Description string `json:"description"`
	ID          string `json:"id"`
	ChildIDs    string `json:"child_ids,omitempty"`
	ParentIDs   string `json:"parent_ids,omitempty"`
	State       string `json:"state"`
	IsDead      bool   `json:"-"`
	Additional  string `json:"additional,omitempty"`
}

func exit(exitCode *int) {
	os.Exit(*exitCode)
}

func TestMain(m *testing.M) {
	var exitCode int = 1
	cmd := startServer()

	defer exit(&exitCode)
	defer attemptRecover(&exitCode)
	defer stopServer(cmd)

	exitCode = m.Run()
	adminPrint(fmt.Sprintf("Test exit code: %d", exitCode))
}

func TestThings(t *testing.T) {
	req := APICall{
		URL:    "http://localhost:8080/things?id=1",
		Method: GET,
	}
	res := req.fire()
	defer res.Body.Close()

	var thing Thing
	err := json.NewDecoder(res.Body).Decode(&thing)
	require.Nil(t, err)

	assert.Equal(t, thing.ID, "1")
	assert.NotEmpty(t, thing.Description)
	assert.NotEmpty(t, thing.State)
}

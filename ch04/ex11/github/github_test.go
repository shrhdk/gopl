package github

import (
	"encoding/json"
	"testing"
)

func TestCreateIssueParam(t *testing.T) {
	var given createIssueParam
	bytes, _ := json.Marshal(given)

	expected := `{"title":"","body":"","assignee":"","milestone":0}`
	actual := string(bytes)

	if actual != expected {
		t.Errorf("\ngot  %c\nwant %c", actual, expected)
	}
}

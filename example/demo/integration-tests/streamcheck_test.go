package integrationtests

import (
	"net/http"
	"testing"
)

func TestAPIStatus(t *testing.T) {
	resp, err := http.Get("http://api:9999/users/1/is-streaming")
	if err != nil {
		t.Fatalf("Error when getting status %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Invalid status code %v", resp.StatusCode)
	}
}

package rimindertest

import (
	"testing"

	"github.com/Xalrandion/go-riminder-api"
)

func TestSourceListNoError(t *testing.T) {
	api := riminder.New(GetTestHelper().APIKey)
	resp, err := api.Source().List()

	if err != nil {
		t.Fatalf("Invalid response: %v", err)
	}
	// will fail if team have no source, tests response parsing also.
	if len(resp) == 0 {
		t.Fatal("No source retrieved")
	}
}

func TestSourceListExpError(t *testing.T) {
	api := riminder.New("an obviously invalid key")
	_, err := api.Source().List()

	if err == nil {
		t.Fatalf("A valid response was not expected.")
	}
}

func TestSourceGetNoErrorFID(t *testing.T) {
	api := riminder.New(GetTestHelper().APIKey)
	resp, err := api.Source().Get(map[string]interface{}{"source_id": GetTestHelper().SourceID})

	if err != nil {
		t.Fatalf("Invalid response: %v", err)
	}
	// tests response parsing also (only on this one because is the more complexe one).
	if resp.SourceID != GetTestHelper().SourceID {
		t.Fatalf("Missmatch source ids (exp: '%s' got: '%s'", GetTestHelper().SourceID, resp.SourceID)
	}
}

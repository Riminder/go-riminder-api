package rimindertest

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"sync"

	"github.com/Xalrandion/go-riminder-api/riminder"
)

var instance *TestHelper
var once sync.Once

// TestHelper contains methods and fields to run the tests.
type TestHelper struct {
	APIKey           string
	WebhookKey       string
	ProfileID        string
	ProfileReference string
	SourceID         string
	FilterID         string
	FilterReference  string
	CvTestPath       string
}

// GetTestHelper returns the instance of TestHelper (which is a singleton).
func GetTestHelper() *TestHelper {
	once.Do(func() {
		instance = &TestHelper{
			APIKey:           "ask_874138568ebde822652c3ddf2218333a",
			WebhookKey:       "quoi",
			ProfileID:        "be50e7e57f4ad045ac8c406e7b665d409be9e363",
			ProfileReference: "5279",
			SourceID:         "fe6d7a2aa9125259a5ecf7905154a0396a891c06",
			FilterID:         "050bdaa9cedad3cf2bc2fac3e4e4acb7499d762a",
			FilterReference:  "1234567",
			CvTestPath:       "./test/assets/cv_test01.pdf",
		}
	})
	return instance
}

func (t *TestHelper) genWebhookRequest() string {
	payload := map[string]interface{}{
		"type":    riminder.EventProfileParseSuccess,
		"message": "Parsing success I gess (:",
		"profile": map[string]string{
			"profile_id":        "hey",
			"profile_reference": "hop",
		},
	}

	jsonPayload, _ := json.Marshal(payload)
	hasher := hmac.New(sha256.New, []byte(GetTestHelper().WebhookKey))

	hasher.Write([]byte(jsonPayload))
	b64Sign := base64.StdEncoding.EncodeToString(hasher.Sum(nil))
	b64Payload := base64.StdEncoding.EncodeToString(jsonPayload)
	return b64Sign + "." + b64Payload
}

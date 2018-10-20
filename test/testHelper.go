package rimindertest

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"sync"

	"github.com/Xalrandion/go-riminder-api/riminder/response"

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
			APIKey:           "",
			WebhookKey:       "quoi",
			ProfileID:        "be50e7e57f4ad045ac8c406e7b665d409be9e363",
			ProfileReference: "5279",
			SourceID:         "fe6d7a2aa9125259a5ecf7905154a0396a891c06",
			FilterID:         "050bdaa9cedad3cf2bc2fac3e4e4acb7499d762a",
			FilterReference:  "1234567",
			CvTestPath:       "./assets/cv_test01.png",
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

func (t *TestHelper) genTrainingMetadata() []riminderResponse.TrainingMetadataElem {
	res := []riminderResponse.TrainingMetadataElem{
		riminderResponse.TrainingMetadataElem{
			FilterReference: t.FilterReference,
			Stage:           "NO",
			StageTimestamp:  1536619600,
			Rating:          3,
			RatingTimestamp: 1536619600,
		},
	}
	return res
}

func (t *TestHelper) genProfileData() riminderResponse.ProfileData {
	resp := riminderResponse.ProfileData{
		Name:            "Some Man",
		Email:           "some.man@gmail.com",
		Phone:           "+34 5 65 34 23 44",
		Summary:         "Some Man is some man",
		Location:        "",
		LocationDetails: riminderResponse.ProfileLocDetails{Text: "23 rue pomme, 345433 Arlandolia"},
		Experiences: []struct {
			Start           string                             `json:"start"`
			End             string                             `json:"end"`
			Title           string                             `json:"title"`
			Company         string                             `json:"company"`
			Location        string                             `json:"location"`
			LocationDetails riminderResponse.ProfileLocDetails `json:"location_details"`
			Description     string                             `json:"description"`
		}{
			{
				Start:           "23/34/2017",
				End:             "",
				Title:           "Some title",
				Company:         "Some company",
				Location:        "",
				LocationDetails: riminderResponse.ProfileLocDetails{Text: "234 rue pomme, 345433 Arlandolia"},
				Description:     "Was cool",
			},
		},
		Educations: []struct {
			Start           string                             `json:"start"`
			End             string                             `json:"end"`
			Title           string                             `json:"title"`
			School          string                             `json:"school"`
			LocationDetails riminderResponse.ProfileLocDetails `json:"location_details"`
			Location        string                             `json:"location"`
			Description     string                             `json:"description"`
		}{
			{
				Start:           "23/34/2017",
				End:             "",
				Title:           "Some title",
				School:          "Some company",
				Location:        "",
				LocationDetails: riminderResponse.ProfileLocDetails{Text: "234 rue pomme, 345433 Arlandolia"},
				Description:     "Was cool",
			},
		},
		Skills:    []string{"skill1", "management"},
		Languages: []string{"french"},
		URLs: struct {
			FromResume []string `json:"from_resume"`
			Linkedin   string   `json:"linkedin"`
			Twitter    string   `json:"twitter"`
			Facebook   string   `json:"facebook"`
			Github     string   `json:"github"`
			Picture    string   `json:"picture"`
		}{
			FromResume: nil,
			Linkedin:   "",
			Twitter:    "",
			Facebook:   "",
			Github:     "",
			Picture:    "",
		},
	}
	return resp
}

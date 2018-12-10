package rimindertest

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/Xalrandion/go-riminder-api"
)

func TestProfileAddNoErrorMaxArgs(t *testing.T) {

	api := riminder.New(GetTestHelper().APIKey)

	args := map[string]interface{}{
		"filepath":            GetTestHelper().CvTestPath,
		"source_id":           GetTestHelper().SourceID,
		"training_metadata":   GetTestHelper().genTrainingMetadata(),
		"profile_reference":   strconv.Itoa(rand.Int()) + "zap",
		"timestamp_reception": 1536619600,
	}
	_, err := api.Profile().Add(args)

	if err != nil {
		t.Fatalf("Invalid response: %v", err)
	}
}

func TestProfileAddNoErrorMinArgs(t *testing.T) {

	api := riminder.New(GetTestHelper().APIKey)

	args := map[string]interface{}{
		"filepath":  GetTestHelper().CvTestPath,
		"source_id": GetTestHelper().SourceID,
	}
	_, err := api.Profile().Add(args)

	if err != nil {
		t.Fatalf("Invalid response: %v", err)
	}
}

func TestProfileListNoErrorMinArgs(t *testing.T) {

	api := riminder.New(GetTestHelper().APIKey)

	args := map[string]interface{}{
		"source_ids": []string{GetTestHelper().SourceID},
	}
	_, err := api.Profile().List(args)

	if err != nil {
		t.Fatalf("Invalid response: %v", err)
	}
}

func TestProfileListNoErrorMinArgsWithFilter(t *testing.T) {

	api := riminder.New(GetTestHelper().APIKey)

	args := map[string]interface{}{
		"source_ids": []string{GetTestHelper().SourceID},
		"filter_id":  GetTestHelper().FilterID,
		"stage":      "NEW",
		"rating":     1,
	}
	_, err := api.Profile().List(args)

	if err != nil {
		t.Fatalf("Invalid response: %v", err)
	}
}

func TestProfileListNoErrorMaxArgsWithFilter(t *testing.T) {

	api := riminder.New(GetTestHelper().APIKey)

	args := map[string]interface{}{
		"source_ids":       []string{GetTestHelper().SourceID},
		"filter_reference": GetTestHelper().FilterReference,
		"stage":            "NEW",
		"rating":           1,
		"seniority":        "all",
		"date_start":       "1347209668",
		"date_end":         "1537374495",
		"page":             2,
		"limit":            3,
		"sort_by":          "ranking",
		"order_by":         "asc",
	}
	_, err := api.Profile().List(args)

	if err != nil {
		t.Fatalf("Invalid response: %v", err)
	}
}

func TestProfileListNoErrorMaxArgs(t *testing.T) {

	api := riminder.New(GetTestHelper().APIKey)

	args := map[string]interface{}{
		"source_ids": []string{GetTestHelper().SourceID},
		"seniority":  "all",
		"date_start": "1347209668",
		"date_end":   "1537374495",
		"page":       2,
		"limit":      3,
		"sort_by":    "ranking",
		"order_by":   "asc",
	}
	_, err := api.Profile().List(args)

	if err != nil {
		t.Fatalf("Invalid response: %v", err)
	}
}

func TestProfileDocumentListNoErrorPID(t *testing.T) {

	api := riminder.New(GetTestHelper().APIKey)

	args := map[string]interface{}{
		"source_id":  GetTestHelper().SourceID,
		"profile_id": GetTestHelper().ProfileID,
	}
	_, err := api.Profile().Document().List(args)

	if err != nil {
		t.Fatalf("Invalid response: %v", err)
	}
}

func TestProfileDocumentListNoErrorPRef(t *testing.T) {

	api := riminder.New(GetTestHelper().APIKey)

	args := map[string]interface{}{
		"source_id":         GetTestHelper().SourceID,
		"profile_reference": GetTestHelper().ProfileReference,
	}
	_, err := api.Profile().Document().List(args)

	if err != nil {
		t.Fatalf("Invalid response: %v", err)
	}
}

func TestProfileScoringListNoErrorPID(t *testing.T) {

	api := riminder.New(GetTestHelper().APIKey)

	args := map[string]interface{}{
		"source_id":  GetTestHelper().SourceID,
		"profile_id": GetTestHelper().ProfileID,
	}
	_, err := api.Profile().Scoring().List(args)

	if err != nil {
		t.Fatalf("Invalid response: %v", err)
	}
}

func TestProfileScoringListNoErrorPRef(t *testing.T) {

	api := riminder.New(GetTestHelper().APIKey)

	args := map[string]interface{}{
		"source_id":         GetTestHelper().SourceID,
		"profile_reference": GetTestHelper().ProfileReference,
	}
	_, err := api.Profile().Scoring().List(args)

	if err != nil {
		t.Fatalf("Invalid response: %v", err)
	}
}

func TestProfileParsingGetNoErrorPID(t *testing.T) {

	api := riminder.New(GetTestHelper().APIKey)

	args := map[string]interface{}{
		"source_id":  GetTestHelper().SourceID,
		"profile_id": GetTestHelper().ProfileID,
	}
	_, err := api.Profile().Parsing().Get(args)

	if err != nil {
		t.Fatalf("Invalid response: %v", err)
	}
}

func TestProfileParsingGetNoErrorPRef(t *testing.T) {

	api := riminder.New(GetTestHelper().APIKey)

	args := map[string]interface{}{
		"source_id":         GetTestHelper().SourceID,
		"profile_reference": GetTestHelper().ProfileReference,
	}
	_, err := api.Profile().Parsing().Get(args)

	if err != nil {
		t.Fatalf("Invalid response: %v", err)
	}
}

func TestProfileStageSetNoErrorPID(t *testing.T) {

	api := riminder.New(GetTestHelper().APIKey)

	args := map[string]interface{}{
		"source_id":  GetTestHelper().SourceID,
		"profile_id": GetTestHelper().ProfileID,
		"filter_id":  GetTestHelper().FilterID,
		"stage":      "NO",
	}
	_, err := api.Profile().Stage().Set(args)

	if err != nil {
		t.Fatalf("Invalid response: %v", err)
	}
}

func TestProfileStageSetNoErrorPRef(t *testing.T) {

	api := riminder.New(GetTestHelper().APIKey)

	args := map[string]interface{}{
		"source_id":         GetTestHelper().SourceID,
		"profile_reference": GetTestHelper().ProfileReference,
		"filter_reference":  GetTestHelper().FilterReference,
		"stage":             "NO",
	}
	_, err := api.Profile().Stage().Set(args)

	if err != nil {
		t.Fatalf("Invalid response: %v", err)
	}
}

func TestProfileRatingSetNoErrorPID(t *testing.T) {

	api := riminder.New(GetTestHelper().APIKey)

	args := map[string]interface{}{
		"source_id":  GetTestHelper().SourceID,
		"profile_id": GetTestHelper().ProfileID,
		"filter_id":  GetTestHelper().FilterID,
		"rating":     1,
	}
	_, err := api.Profile().Rating().Set(args)

	if err != nil {
		t.Fatalf("Invalid response: %v", err)
	}
}

func TestProfileRatingSetNoErrorPRef(t *testing.T) {

	api := riminder.New(GetTestHelper().APIKey)

	args := map[string]interface{}{
		"source_id":         GetTestHelper().SourceID,
		"profile_reference": GetTestHelper().ProfileReference,
		"filter_reference":  GetTestHelper().FilterReference,
		"rating":            1,
	}
	_, err := api.Profile().Rating().Set(args)

	if err != nil {
		t.Fatalf("Invalid response: %v", err)
	}
}

func TestProfileJsonAddNoErrorMinArg(t *testing.T) {

	api := riminder.New(GetTestHelper().APIKey)

	args := map[string]interface{}{
		"source_id":    GetTestHelper().SourceID,
		"profile_json": GetTestHelper().genProfileData(),
	}
	_, err := api.Profile().JSON().Add(args)

	if err != nil {
		t.Fatalf("Invalid response: %v", err)
	}
}

func TestProfileJsonAddNoErrorMaxArg(t *testing.T) {

	api := riminder.New(GetTestHelper().APIKey)

	args := map[string]interface{}{
		"source_id":           GetTestHelper().SourceID,
		"profile_json":        GetTestHelper().genProfileData(),
		"training_metadata":   GetTestHelper().genTrainingMetadata(),
		"profile_reference":   strconv.Itoa(rand.Int()) + "zop",
		"timestamp_reception": 1536619600,
	}
	_, err := api.Profile().JSON().Add(args)

	if err != nil {
		t.Fatalf("Invalid response: %v", err)
	}
}

func TestProfileJsonCheckNoErrorMinArg(t *testing.T) {

	api := riminder.New(GetTestHelper().APIKey)

	args := map[string]interface{}{
		"profile_json": GetTestHelper().genProfileData(),
	}
	_, err := api.Profile().JSON().Check(args)

	if err != nil {
		t.Fatalf("Invalid response: %v", err)
	}
}

func TestProfileJsonCheckNoErrorMaxArg(t *testing.T) {

	api := riminder.New(GetTestHelper().APIKey)

	args := map[string]interface{}{
		"profile_json": GetTestHelper().genProfileData(),
	}
	_, err := api.Profile().JSON().Check(args)

	if err != nil {
		t.Fatalf("Invalid response: %v", err)
	}
}

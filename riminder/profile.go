package riminder

import (
	"encoding/json"
	"fmt"

	"github.com/Xalrandion/go-riminder-api/riminder/response"
)

// profile class contains methods to interact with the profiles on riminder's api.
type profile struct {
	client    *clientw
	Documents *profileDocument
	Parsing   *profileParsing
	Scoring   *profileScoring
	Stage     *profileStage
	Rating    *profileRating
	JSON      *profileStructData
}

func newprofile(riminder *Riminder) *profile {
	s := new(profile)

	s.client = riminder.clientw
	s.Documents = newprofileDocument(riminder)
	s.Parsing = newprofileParsing(riminder)
	s.Scoring = newprofileScoring(riminder)
	s.Stage = newprofileStage(riminder)
	s.Rating = newprofileRating(riminder)
	s.JSON = newprofileStructData(riminder)
	return s
}

func addTrainingMetadataStringIfNotEmpty(params *map[string]string, options map[string]interface{}) {
	if tMetadata, ok := options["training_metadata"]; ok {
		btMetadata, err := json.Marshal(tMetadata)
		if err != nil {
			(*params)["training_metadata"] = string(btMetadata)
		}
	}
}

// Add add a new profile to a source.
func (p *profile) Add(options map[string]interface{}) (riminderResponse.ProfileAddElem, error) {
	filepath := options["filepath"].(string)

	params := map[string]string{
		"source_id": options["source_id"].(string),
	}
	addTrainingMetadataStringIfNotEmpty(&params, options)
	AddIfNotEmptyStrMap(&params, options, "profile_reference")
	AddIfNotEmptyStrMap(&params, options, "timestamp_reception")

	resp := riminderResponse.ProfileAddContainer{}
	err := p.client.PostFile("profile", filepath, params, &resp)
	if err != nil {
		return riminderResponse.ProfileAddElem{}, err
	}
	return resp.Data, nil
}

// List get a list of profiles.
func (p *profile) List(options map[string]interface{}) (riminderResponse.ProfileListElem, error) {
	sourceIDs, err := json.Marshal(options["source_id"])
	if err != nil {
		return riminderResponse.ProfileListElem{}, fmt.Errorf("profile.List:Cannot parse source ids (should be a list of string): %v", err)
	}

	query := map[string]string{
		"source_ids": string(sourceIDs),
	}
	AddIfNotEmptyStrMap(&query, options, "seniority")
	AddIfNotEmptyStrMap(&query, options, "filter_id")
	AddIfNotEmptyStrMap(&query, options, "filter_reference")
	AddIfNotEmptyStrMap(&query, options, "stage")
	AddIfNotEmptyStrMap(&query, options, "rating")
	AddIfNotEmptyStrMap(&query, options, "date_start")
	AddIfNotEmptyStrMap(&query, options, "date_end")
	AddIfNotEmptyStrMap(&query, options, "page")
	AddIfNotEmptyStrMap(&query, options, "limit")
	AddIfNotEmptyStrMap(&query, options, "sort_by")
	AddIfNotEmptyStrMap(&query, options, "order_by")

	resp := riminderResponse.ProfileListContainer{}
	err = p.client.Get("profiles", query, &resp)
	if err != nil {
		return riminderResponse.ProfileListElem{}, err
	}
	return resp.Data, nil
}

// Get a specific profile.
func (p *profile) Get(options map[string]interface{}) (riminderResponse.ProfileGetElem, error) {
	query := map[string]string{
		"source_id": options["source_id"].(string),
	}
	AddIfNotEmptyStrMap(&query, options, "profile_id")
	AddIfNotEmptyStrMap(&query, options, "profile_reference")

	resp := riminderResponse.ProfileGetContainer{}
	err := p.client.Get("profile", query, &resp)
	if err != nil {
		return riminderResponse.ProfileGetElem{}, err
	}
	return resp.Data, nil
}

type profileDocument struct {
	client *clientw
}

func newprofileDocument(riminder *Riminder) *profileDocument {
	s := new(profileDocument)

	s.client = riminder.clientw
	return s
}

// List get the list of attachement for a profile.
func (p *profileDocument) List(options map[string]interface{}) ([]riminderResponse.ProfileDocumentsListElem, error) {
	query := map[string]string{
		"source_id": options["source_id"].(string),
	}
	AddIfNotEmptyStrMap(&query, options, "profile_id")
	AddIfNotEmptyStrMap(&query, options, "profile_reference")

	resp := riminderResponse.ProfileDocumentsListContainer{}
	err := p.client.Get("profile/documents", query, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

type profileParsing struct {
	client *clientw
}

func newprofileParsing(riminder *Riminder) *profileParsing {
	s := new(profileParsing)

	s.client = riminder.clientw
	return s
}

// Get gets parsing infos for a specific profile.
func (p *profileParsing) Get(options map[string]interface{}) (riminderResponse.ProfileParsingGetElem, error) {
	query := map[string]string{
		"source_id": options["source_id"].(string),
	}
	AddIfNotEmptyStrMap(&query, options, "profile_id")
	AddIfNotEmptyStrMap(&query, options, "profile_reference")

	resp := riminderResponse.ProfileParsingGetContainer{}
	err := p.client.Get("profile/parsing", query, &resp)
	if err != nil {
		return riminderResponse.ProfileParsingGetElem{}, err
	}
	return resp.Data, nil
}

type profileScoring struct {
	client *clientw
}

func newprofileScoring(riminder *Riminder) *profileScoring {
	s := new(profileScoring)

	s.client = riminder.clientw
	return s
}

// Get gets scoring infos for a specific profile.
func (p *profileScoring) List(options map[string]interface{}) ([]riminderResponse.ProfileScoringListElem, error) {
	query := map[string]string{
		"source_id": options["source_id"].(string),
	}
	AddIfNotEmptyStrMap(&query, options, "profile_id")
	AddIfNotEmptyStrMap(&query, options, "profile_reference")

	resp := riminderResponse.ProfileScoringListContainer{}
	err := p.client.Get("profile/scoring", query, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

type profileStage struct {
	client *clientw
}

func newprofileStage(riminder *Riminder) *profileStage {
	s := new(profileStage)

	s.client = riminder.clientw
	return s
}

// Set sets the stage for a specific profile and filter.
func (p *profileStage) Set(options map[string]interface{}) (riminderResponse.ProfileStageSetElem, error) {

	resp := riminderResponse.ProfileStageSetContainer{}
	err := p.client.Patch("profile/stage", options, &resp)
	if err != nil {
		return riminderResponse.ProfileStageSetElem{}, err
	}
	return resp.Data, nil
}

type profileRating struct {
	client *clientw
}

func newprofileRating(riminder *Riminder) *profileRating {
	s := new(profileRating)

	s.client = riminder.clientw
	return s
}

// Set sets the rating for a specific profile and filter.
func (p *profileRating) Set(options map[string]interface{}) (riminderResponse.ProfileRatingSetElem, error) {

	resp := riminderResponse.ProfileRatingSetContainer{}
	err := p.client.Patch("profile/rating", options, &resp)
	if err != nil {
		return riminderResponse.ProfileRatingSetElem{}, err
	}
	return resp.Data, nil
}

type profileStructData struct {
	client *clientw
}

func newprofileStructData(riminder *Riminder) *profileStructData {
	s := new(profileStructData)

	s.client = riminder.clientw
	return s
}

// Check checks if the given parsed profile is valid.
func (s *profileStructData) Check(options map[string]interface{}) (riminderResponse.ProfileJSONCheckElem, error) {
	resp := riminderResponse.ProfileJSONCheckContainer{}
	err := s.client.Post("profile/json/check", options, &resp)
	if err != nil {
		return riminderResponse.ProfileJSONCheckElem{}, err
	}
	return resp.Data, nil
}

// Add adds a new parsed profile.
func (s *profileStructData) Add(options map[string]interface{}) (riminderResponse.ProfileJSONAddElem, error) {
	resp := riminderResponse.ProfileJSONAddContainer{}

	err := s.client.Post("profile/json", options, &resp)
	if err != nil {
		return riminderResponse.ProfileJSONAddElem{}, err
	}
	return resp.Data, nil
}

package riminder

import (
	"encoding/json"
	"fmt"

	"github.com/Xalrandion/go-riminder-api/riminder/response"
)

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

func (p *profile) Add(options map[string]interface{}) (response.ProfileAddElem, error) {
	filepath := options["filepath"].(string)

	params := map[string]string{
		"source_id": options["source_id"].(string),
	}
	addTrainingMetadataStringIfNotEmpty(&params, options)
	AddIfNotEmptyStrMap(&params, options, "profile_reference")
	AddIfNotEmptyStrMap(&params, options, "timestamp_reception")

	resp := response.ProfileAddContainer{}
	err := p.client.PostFile("profile", filepath, params, &resp)
	if err != nil {
		return response.ProfileAddElem{}, err
	}
	return resp.Data, nil
}

func (p *profile) List(options map[string]interface{}) (response.ProfileListElem, error) {
	sourceIDs, err := json.Marshal(options["source_id"])
	if err != nil {
		return response.ProfileListElem{}, fmt.Errorf("profile.List:Cannot parse source ids (should be a list of string): %v", err)
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

	resp := response.ProfileListContainer{}
	err = p.client.Get("profiles", query, &resp)
	if err != nil {
		return response.ProfileListElem{}, err
	}
	return resp.Data, nil
}

func (p *profile) Get(options map[string]interface{}) (response.ProfileGetElem, error) {
	query := map[string]string{
		"source_id": options["source_id"].(string),
	}
	AddIfNotEmptyStrMap(&query, options, "profile_id")
	AddIfNotEmptyStrMap(&query, options, "profile_reference")

	resp := response.ProfileGetContainer{}
	err := p.client.Get("profile", query, &resp)
	if err != nil {
		return response.ProfileGetElem{}, err
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

func (p *profileDocument) List(options map[string]interface{}) ([]response.ProfileDocumentsListElem, error) {
	query := map[string]string{
		"source_id": options["source_id"].(string),
	}
	AddIfNotEmptyStrMap(&query, options, "profile_id")
	AddIfNotEmptyStrMap(&query, options, "profile_reference")

	resp := response.ProfileDocumentsListContainer{}
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

func (p *profileParsing) Get(options map[string]interface{}) (response.ProfileParsingGetElem, error) {
	query := map[string]string{
		"source_id": options["source_id"].(string),
	}
	AddIfNotEmptyStrMap(&query, options, "profile_id")
	AddIfNotEmptyStrMap(&query, options, "profile_reference")

	resp := response.ProfileParsingGetContainer{}
	err := p.client.Get("profile/parsing", query, &resp)
	if err != nil {
		return response.ProfileParsingGetElem{}, err
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

func (p *profileScoring) List(options map[string]interface{}) ([]response.ProfileScoringListElem, error) {
	query := map[string]string{
		"source_id": options["source_id"].(string),
	}
	AddIfNotEmptyStrMap(&query, options, "profile_id")
	AddIfNotEmptyStrMap(&query, options, "profile_reference")

	resp := response.ProfileScoringListContainer{}
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

func (p *profileStage) Set(options map[string]interface{}) (response.ProfileStageSetElem, error) {

	resp := response.ProfileStageSetContainer{}
	err := p.client.Patch("profile/stage", options, &resp)
	if err != nil {
		return response.ProfileStageSetElem{}, err
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

func (p *profileRating) Set(options map[string]interface{}) (response.ProfileRatingSetElem, error) {

	resp := response.ProfileRatingSetContainer{}
	err := p.client.Patch("profile/rating", options, &resp)
	if err != nil {
		return response.ProfileRatingSetElem{}, err
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

func (s *profileStructData) Check(options map[string]interface{}) (response.ProfileJSONCheckElem, error) {
	resp := response.ProfileJSONCheckContainer{}
	err := s.client.Post("profile/json/check", options, &resp)
	if err != nil {
		return response.ProfileJSONCheckElem{}, err
	}
	return resp.Data, nil
}

func (s *profileStructData) Add(options map[string]interface{}) (response.ProfileJSONAddElem, error) {
	resp := response.ProfileJSONAddContainer{}

	err := s.client.Post("profile/json", options, &resp)
	if err != nil {
		return response.ProfileJSONAddElem{}, err
	}
	return resp.Data, nil
}

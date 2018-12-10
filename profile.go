package riminder

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/Xalrandion/go-riminder-api/response"
)

// defaultDateStart default timestamp for proofile.List
const defaultDateStart = "1347209668"

// profile class contains methods to interact with the profiles on riminder's api.
type profile struct {
	client     *clientw
	documents  *profileDocument
	parsing    *profileParsing
	scoring    *profileScoring
	revealing  *profileRevealing
	stage      *profileStage
	rating     *profileRating
	structured *profileStructData
}

func newprofile(riminder *Riminder) *profile {
	s := new(profile)

	s.client = riminder.clientw
	s.documents = newprofileDocument(riminder)
	s.parsing = newprofileParsing(riminder)
	s.scoring = newprofileScoring(riminder)
	s.revealing = newprofileRevealing(riminder)
	s.stage = newprofileStage(riminder)
	s.rating = newprofileRating(riminder)
	s.structured = newprofileStructData(riminder)
	return s
}

func (p *profile) Document() *profileDocument {
	return p.documents
}

func (p *profile) Parsing() *profileParsing {
	return p.parsing
}

func (p *profile) Scoring() *profileScoring {
	return p.scoring
}

func (p *profile) Revealing() *profileRevealing {
	return p.revealing
}

func (p *profile) Stage() *profileStage {
	return p.stage
}

func (p *profile) Rating() *profileRating {
	return p.rating
}

func (p *profile) JSON() *profileStructData {
	return p.structured
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
	sourceIDs, err := json.Marshal(options["source_ids"])
	if err != nil {
		return riminderResponse.ProfileListElem{}, fmt.Errorf("profile.List:Cannot parse source ids (should be a list of string): %v", err)
	}

	query := map[string]string{
		"source_ids": string(sourceIDs),
		"sort_by":    "ranking",
		"page":       "1",
		"date_end":   strconv.Itoa(int(time.Now().Unix())),
		"date_start": defaultDateStart,
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

// Get scoring infos for a specific profile.
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

type profileRevealing struct {
	client *clientw
}

func newprofileRevealing(riminder *Riminder) *profileRevealing {
	s := new(profileRevealing)

	s.client = riminder.clientw
	return s
}

// Get gets revealing infos for a specific profile.
func (p *profileRevealing) Get(options map[string]interface{}) ([]riminderResponse.ProfileRevealingGetElem, error) {
	query := map[string]string{
		"source_id": options["source_id"].(string),
	}
	AddIfNotEmptyStrMap(&query, options, "profile_id")
	AddIfNotEmptyStrMap(&query, options, "profile_reference")
	AddIfNotEmptyStrMap(&query, options, "filter_id")
	AddIfNotEmptyStrMap(&query, options, "filter_reference")

	resp := riminderResponse.ProfileRevealingGetContainer{}
	err := p.client.Get("profile/revealing", query, &resp)
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

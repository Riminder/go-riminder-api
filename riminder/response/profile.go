package response

// ProfileListContainer contains the response representation for profile.list method.
type ProfileListContainer struct {
	*Container

	Data ProfileListElem `json:"data"`
}

// ProfileListElem contains the data field response representation for profile.list method.
type ProfileListElem struct {
	Page          int `json:"page"`
	MaxPage       int `json:"maxPage"`
	CountProfiles int `json:"count_profiles"`
	Profiles      []struct {
		ProfileID        string   `json:"profile_id"`
		ProfileReference string   `json:"profile_reference"`
		Name             string   `json:"name"`
		Email            string   `json:"email"`
		Seniority        string   `json:"seniority"`
		DateReception    DateTime `json:"date_reception"`
		DateCreation     DateTime `json:"date_creation"`
		Source           struct {
			SourceID string `json:"source_id"`
			Name     string `json:"name"`
			Type     string `json:"type"`
		} `json:"source"`
		Score  float64 `json:"score"`
		Stage  string  `json:"stage"`
		Rating int     `json:"rating"`
	} `json:"profiles"`
}

// ProfileAddContainer contains the response representation for profile.add method.
type ProfileAddContainer struct {
	*Container

	Data ProfileAddElem
}

// ProfileAddElem contains the data field response representation for profile.add method.
type ProfileAddElem struct {
	ProfileReference string   `json:"profile_reference"`
	FileID           string   `json:"file_id"`
	FileName         string   `json:"file_name"`
	FileSize         string   `json:"file_size"`
	Extension        string   `json:"extension"`
	DateReception    DateTime `json:"date_reception"`
}

// ProfileLocDetails a specific field in responses.
type ProfileLocDetails struct {
	Text string `json:"text"`
}

// ProfileData a specific field in responses, usable as request for profile.json*.
type ProfileData struct {
	Name            string            `json:"name"`
	Email           string            `json:"email"`
	Phone           string            `json:"phone"`
	Summary         string            `json:"summary"`
	Location        string            `json:"location"`
	LocationDetails ProfileLocDetails `json:"location_details"`
	Experiences     []struct {
		Start           string            `json:"start"`
		End             string            `json:"end"`
		Title           string            `json:"title"`
		Company         string            `json:"company"`
		Location        string            `json:"location"`
		LocationDetails ProfileLocDetails `json:"location_details"`
		Description     string            `json:"description"`
	} `json:"experience"`
	Educations []struct {
		Start           string            `json:"start"`
		End             string            `json:"end"`
		Title           string            `json:"title"`
		School          string            `json:"school"`
		LocationDetails ProfileLocDetails `json:"location_details"`
		Location        string            `json:"location"`
		Description     string            `json:"description"`
	} `json:"educations"`
	Skills    []string `json:"skills"`
	Languages []string `json:"languages"`
	Interests []string `json:"interests"`
	URLs      struct {
		FromResume []string `json:"from_resume"`
		Linkedin   string   `json:"linkedin"`
		Twitter    string   `json:"twitter"`
		Facebook   string   `json:"facebook"`
		Github     string   `json:"github"`
		Picture    string   `json:"picture"`
	} `json:"urls"`
}

// TrainingMetadataElem a specific field in responses, usable as request for profile.json*.
type TrainingMetadataElem struct {
	FilterReference string `json:"filter_reference"`
	Stage           string `json:"stage"`
	StageTimestamp  int    `json:"stage_timestamp"`
	Rating          int    `json:"rating"`
	RatingTimestamp int    `json:"rating_timestamp"`
}

// ProfileJSONCheckContainer contains the response representation for profile.json.check method.
type ProfileJSONCheckContainer struct {
	*Container

	Data ProfileJSONCheckElem `json:"data"`
}

// ProfileJSONCheckElem contains the data field response representation for profile.json.check method.
type ProfileJSONCheckElem struct {
	ProfileJSON      ProfileData            `json:"profile_json"`
	TrainingMetadata []TrainingMetadataElem `json:"training_metadata"`
}

// ProfileJSONAddContainer contains the response representation for profile.json.add method.
type ProfileJSONAddContainer struct {
	*Container

	Data ProfileJSONAddElem `json:"data"`
}

// ProfileJSONAddElem contains the data field response representation for profile.json.add method.
type ProfileJSONAddElem struct {
	ProfileJSON      ProfileData            `json:"profile_json"`
	TrainingMetadata []TrainingMetadataElem `json:"training_metadata"`
	SourceID         string                 `json:"source_id"`
	ProfileID        string                 `json:"profile_id"`
}

// ProfileGetContainer contains the response representation for profile.get method.
type ProfileGetContainer struct {
	*Container

	Data ProfileGetElem `json:"data"`
}

// ProfileGetElem contains the data field response representation for profile.get method.
type ProfileGetElem struct {
	ProfileID        string   `json:"profile_id"`
	ProfileReference string   `json:"profile_reference"`
	Name             string   `json:"name"`
	Email            string   `json:"email"`
	Phone            string   `json:"phone"`
	Address          string   `json:"address"`
	DateReception    DateTime `json:"date_reception"`
	DateCreation     DateTime `json:"date_creation"`
	SourceID         string   `json:"source_id"`
}

// ProfileDocumentsListContainer contains the response representation for profile.document.List method.
type ProfileDocumentsListContainer struct {
	*Container

	Data []ProfileDocumentsListElem `json:"data"`
}

// ProfileDocumentsListElem the data field response representation for profile.document.List method.
type ProfileDocumentsListElem struct {
	Type             string `json:"type"`
	FileName         string `json:"file_name"`
	OriginalFileName string `json:"original_file_name"`
	FileSize         string `json:"file_size"`
	Extension        string `json:"extension"`
	URL              string `json:"url"`
	Timestamp        int    `json:"timestamp"`
	SourceID         string `json:"source_id"`
}

// ProfileParsingGetContainer contains the response representation for profile.parsing.Get method.
type ProfileParsingGetContainer struct {
	*Container

	Data ProfileParsingGetElem `json:"data"`
}

// ProfileParsingGetElem the data field response representation for profile.parsing.Get method.
type ProfileParsingGetElem struct {
	HardSkills  []string `json:"hard_skills"`
	SoftSkills  []string `json:"soft_skills"`
	Languages   []string `json:"languages"`
	Seniority   string   `json:"seniority"`
	Experiences []struct {
		Title       string   `json:"title"`
		Description string   `json:"description"`
		Company     string   `json:"company"`
		Location    string   `json:"location"`
		StartDate   DateTime `json:"start_date"`
		EndDate     DateTime `json:"end_date"`
	} `json:"experiences"`
	Educations []struct {
		Title       string   `json:"title"`
		Description string   `json:"description"`
		School      string   `json:"school"`
		Location    string   `json:"location"`
		StartDate   DateTime `json:"start_date"`
		EndDate     DateTime `json:"end_date"`
	} `json:"educations"`
	DateReception DateTime `json:"date_reception"`
	DateCreation  DateTime `json:"date_creation"`
	SourceID      string   `json:"source_id"`
}

// ProfileScoringListContainer contains the response representation for profile.scoring.List method.
type ProfileScoringListContainer struct {
	*Container

	Data []ProfileScoringListElem `json:"data"`
}

// ProfileScoringListElem the data field response representation for profile.scoring.List method.
type ProfileScoringListElem struct {
	FilterID        string         `json:"filter_id"`
	FilterReference string         `json:"filter_reference"`
	Template        FilterTemplate `json:"template"`
	Score           float64        `json:"score"`
	Stage           string         `json:"stage"`
	Rating          int            `json:"rating"`
	PrivateURL      string         `json:"private_url"`
}

// ProfileStageSetContainer contains the response representation for profile.stage.Set method.
type ProfileStageSetContainer struct {
	*Container

	Data ProfileStageSetElem `json:"data"`
}

// ProfileStageSetElem the data field response representation for profile.stage.Set method.
type ProfileStageSetElem struct {
	FilterID         string `json:"filter_id"`
	FilterReference  string `json:"filter_reference"`
	ProfileID        string `json:"profile_id"`
	ProfileReference string `json:"profile_reference"`
	Stage            string `json:"stage"`
}

// ProfileRatingSetContainer contains the response representation for profile.rating.Set method.
type ProfileRatingSetContainer struct {
	*Container

	Data ProfileRatingSetElem `json:"data"`
}

// ProfileRatingSetElem the data field response representation for profile.rating.Set method.
type ProfileRatingSetElem struct {
	FilterID         string `json:"filter_id"`
	FilterReference  string `json:"filter_reference"`
	ProfileID        string `json:"profile_id"`
	ProfileReference string `json:"profile_reference"`
	Rating           int    `json:"rating"`
}

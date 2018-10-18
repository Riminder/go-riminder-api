package response

type ProfileListContainer struct {
	*ResponseContainer

	Data ProfileListElem `json:"data"`
}

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

type ProfileAddContainer struct {
	*ResponseContainer

	Data ProfileAddElem
}

type ProfileAddElem struct {
	ProfileReference string   `json:"profile_reference"`
	FileID           string   `json:"file_id"`
	FileName         string   `json:"file_name"`
	FileSize         string   `json:"file_size"`
	Extension        string   `json:"extension"`
	DateReception    DateTime `json:"date_reception"`
}

type LocDetails struct {
	Text string `json:"text"`
}

type ProfileData struct {
	Name            string     `json:"name"`
	Email           string     `json:"email"`
	Phone           string     `json:"phone"`
	Summary         string     `json:"summary"`
	Location        string     `json:"location"`
	LocationDetails LocDetails `json:"location_details"`
	Experiences     []struct {
		Start           string     `json:"start"`
		End             string     `json:"end"`
		Title           string     `json:"title"`
		Company         string     `json:"company"`
		Location        string     `json:"location"`
		LocationDetails LocDetails `json:"location_details"`
		Description     string     `json:"description"`
	} `json:"experience"`
	Educations []struct {
		Start           string     `json:"start"`
		End             string     `json:"end"`
		Title           string     `json:"title"`
		School          string     `json:"school"`
		LocationDetails LocDetails `json:"location_details"`
		Location        string     `json:"location"`
		Description     string     `json:"description"`
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

type TrainingMetadataElem struct {
	FilterReference string `json:"filter_reference"`
	Stage           string `json:"stage"`
	StageTimestamp  int    `json:"stage_timestamp"`
	Rating          int    `json:"rating"`
	RatingTimestamp int    `json:"rating_timestamp"`
}

type ProfileJSONCheckContainer struct {
	*ResponseContainer

	Data ProfileJSONCheckElem `json:"data"`
}

type ProfileJSONCheckElem struct {
	ProfileJSON      ProfileData            `json:"profile_json"`
	TrainingMetadata []TrainingMetadataElem `json:"training_metadata"`
}

type ProfileJSONAddContainer struct {
	*ResponseContainer

	Data ProfileJSONAddElem `json:"data"`
}

type ProfileJSONAddElem struct {
	ProfileJSON      ProfileData            `json:"profile_json"`
	TrainingMetadata []TrainingMetadataElem `json:"training_metadata"`
	SourceID         string                 `json:"source_id"`
	ProfileID        string                 `json:"profile_id"`
}

type ProfileGetContainer struct {
	*ResponseContainer

	Data ProfileGetElem `json:"data"`
}

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

type ProfileDocumentsListContainer struct {
	*ResponseContainer

	Data []ProfileDocumentsListElem `json:"data"`
}

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

type ProfileParsingGetContainer struct {
	*ResponseContainer

	Data ProfileParsingGetElem `json:"data"`
}

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

type ProfileScoringListContainer struct {
	*ResponseContainer

	Data []ProfileScoringListElem `json:"data"`
}

type ProfileScoringListElem struct {
	FilterID        string         `json:"filter_id"`
	FilterReference string         `json:"filter_reference"`
	Template        FilterTemplate `json:"template"`
	Score           float64        `json:"score"`
	Stage           string         `json:"stage"`
	Rating          int            `json:"rating"`
	PrivateURL      string         `json:"private_url"`
}

type ProfileStageContainer struct {
	*ResponseContainer

	Data ProfileStageElem `json:"data"`
}

type ProfileStageElem struct {
	FilterID         string `json:"filter_id"`
	FilterReference  string `json:"filter_reference"`
	ProfileID        string `json:"profile_id"`
	ProfileReference string `json:"profile_reference"`
	Stage            string `json:"stage"`
}

type ProfileRatingContainer struct {
	*ResponseContainer

	Data ProfileRatingElem `json:"data"`
}

type ProfileRatingElem struct {
	FilterID         string `json:"filter_id"`
	FilterReference  string `json:"filter_reference"`
	ProfileID        string `json:"profile_id"`
	ProfileReference string `json:"profile_reference"`
	Rating           int    `json:"rating"`
}

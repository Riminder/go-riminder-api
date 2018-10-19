package riminderResponse

// FilterListContainer contains the response representation for filter.list method.
type FilterListContainer struct {
	*Container

	Data []FilterListElem
}

// FilterListElem contains the data field response representation for filter.list method.
type FilterListElem struct {
	FilterID        string   `json:"filter_id"`
	FilterReference string   `json:"filter_reference"`
	Name            string   `json:"name"`
	Archive         bool     `json:"archive"`
	DateCreation    DateTime `json:"date_creation"`
}

// FilterGetContainer contains the response representation for filter.get method.
type FilterGetContainer struct {
	*Container

	Data FilterGetElem
}

// FilterTemplate a specific field in responses.
type FilterTemplate struct {
	Name string `json:"name"`
}

// FilterGetElem contains the data field response representation for filter.get method.
type FilterGetElem struct {
	FilterID        string         `json:"filter_id"`
	FilterReference string         `json:"filter_reference"`
	Name            string         `json:"name"`
	Description     string         `json:"description"`
	Template        FilterTemplate `json:"template"`
	Seniority       string         `json:"seniority"`
	Skills          []string       `json:"skills"`
	Stages          struct {
		CountYes   int `json:"count_yes"`
		CountLater int `json:"count_later"`
		CountNo    int `json:"count_no"`
	} `json:"stages"`
	ScoreTreshold int      `json:"score_threshold"`
	Archive       bool     `json:"archive"`
	DateCreation  DateTime `json:"date_creation"`
}

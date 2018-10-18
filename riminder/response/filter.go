package response

type FilterListContainer struct {
	*ResponseContainer

	Data []FilterListElem
}

type FilterListElem struct {
	FilterID        string   `json:"filter_id"`
	FilterReference string   `json:"filter_reference"`
	Name            string   `json:"name"`
	Archive         bool     `json:"archive"`
	DateCreation    DateTime `json:"date_creation"`
}

type FilterGetContainer struct {
	*ResponseContainer

	Data FilterGetElem
}

type FilterTemplate struct {
	Name string `json:"name"`
}

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
	ScoreTreshold string   `json:"score_threshold"`
	Archive       bool     `json:"archive"`
	DateCreation  DateTime `json:"date_creation"`
}

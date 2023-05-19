package scraper

type privacy struct {
	Data []privacyDatum `json:"data"`
}

type privacyDatum struct {
	Attributes attributes `json:"attributes"`
}

type attributes struct {
	PrivacyDetails privacyDetails `json:"privacyDetails"`
}

type privacyDetails struct {
	ManagePrivacyChoicesURL any           `json:"managePrivacyChoicesUrl"`
	PrivacyTypes            []privacyType `json:"privacyTypes"`
}

type privacyType struct {
	DataCategories []dataCategory `json:"dataCategories"`
	Description    string         `json:"description"`
	Identifier     string         `json:"identifier"`
	PrivacyType    string         `json:"privacyType"`
	Purposes       []purpose      `json:"purposes"`
}

type dataCategory struct {
	DataCategory string   `json:"dataCategory"`
	DataTypes    []string `json:"dataTypes"`
	Identifier   string   `json:"identifier"`
}

type purpose struct {
	DataCategories []dataCategory `json:"dataCategories"`
	Identifier     string         `json:"identifier"`
	Purpose        string         `json:"purpose"`
}

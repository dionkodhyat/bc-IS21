package models

// InputData is a struct representation of the expected input received from the input topic
type InputData struct {
	ID                         string `json:"id"`
	NumberOfChildren           int    `json:"numberOfChildren"`
	FamilyComposition          string `json:"familyComposition"`
	FamilyUnitInPayForDecember bool   `json:"familyUnitInPayForDecember"`
}

// OutputData is a struct representation of the data that will be published to the output topic
type OutputData struct {
	ID               string  `json:"id"`
	IsEligible       bool    `json:"isEligible"`
	BaseAmount       float64 `json:"baseAmount"`
	ChildrenAmount   float64 `json:"childrenAmount"`
	SupplementAmount float64 `json:"supplementAmount"`
}

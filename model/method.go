package model

type ParametersInput struct {
	Parameter1 string `json:"parameter1"`
	Parameter2 string `json:"parameter2"`
	Parameter3 string `json:"parameter3"`
}

type ParametersUpdate struct {
	Parameter1 string  `json:"parameter1"`
	Parameter2 *string `json:"parameter2"`
	Parameter3 *string `json:"parameter3"`
}

type InfoResponse struct {
	ID         string `json:"id"`
	Parameter1 string `json:"parameter1"`
	Parameter2 string `json:"parameter2"`
	Parameter3 string `json:"paramerter3"`
}

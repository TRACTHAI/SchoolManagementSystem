package model

// type ParametersInput struct {
// 	Parameter1 string `json:"parameter1"`
// 	Parameter2 string `json:"parameter2"`
// 	Parameter3 string `json:"parameter3"`
// }

type ParametersInput struct { // student
	Gender      string `json:"gender"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	BirthDate   string `json:"birthdate"`
	Nationality string `json:"nationality"`
	Ethnicity   string `json:"ethnicity"`
}

// type ParametersUpdate struct {
// 	Parameter1 string  `json:"parameter1"`
// 	Parameter2 *string `json:"parameter2"`
// 	Parameter3 *string `json:"parameter3"`
// }

type ParametersUpdate struct {
	StudentID   string  `json:"studentID"`
	Gender      *string `json:"gender"`
	Firstname   *string `json:"firstname"`
	Lastname    *string `json:"lastname"`
	BirthDate   *string `json:"birthdate"`
	Nationality *string `json:"nationality"`
	Ethnicity   *string `json:"ethnicity"`
}

type InfoResponse struct {
	StudentID   string  `json:"studentID"`
	Gender      *string `json:"gender"`
	Firstname   *string `json:"firstname"`
	Lastname    *string `json:"lastname"`
	BirthDate   *string `json:"birthdate"`
	Nationality *string `json:"nationality"`
	Ethnicity   *string `json:"ethnicity"`
}

type InfoResponseAll struct {
	StudentID   string  `json:"studentID"`
	Gender      *string `json:"gender"`
	Firstname   *string `json:"firstname"`
	Lastname    *string `json:"lastname"`
	BirthDate   *string `json:"birthdate"`
	Nationality *string `json:"nationality"`
	Ethnicity   *string `json:"ethnicity"`
}

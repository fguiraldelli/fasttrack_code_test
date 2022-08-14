package model

// user represents data about the users and theirs respective answers
type Registred_user struct {
	ID                       string     `json:"id"`
	Name                     string     `json:"name"`
	Email                    string     `json:"email"`
	Quiz                     []Question `json:"questions"`
	Number_corrected_answers int        `json:"number_corrected_answers"`
	User_rated               float64    `json:"user_rated"`
}
package model

// question represents data about a question and its answer
type Question struct {
	ID             string   `json:"id"`
	Question       string   `json:"question"`
	Answers        []string `json:"answers"`
	Correct_answer string   `json:"correct_answer"`
	Answered       bool     `json:"answered"`
	Is_corrected   bool     `json:"is_corrected"`
}
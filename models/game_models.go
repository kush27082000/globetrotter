package models

type User struct {
	Username string `gorm:"unique;not null"`
}

type DestinationDetails struct {
	ID            uint   `gorm:"primaryKey"`
	City          string `json:"city"`
	Country       string `json:"country"`
	Clues         string `json:"clues"`
	FunFact       string `json:"fun_fact"`
	Trivia        string `json:"trivia"`
	CorrectAnswer string `json:"correct_answer"`
}

type ClueDetails struct {
	ID    uint   `gorm:"primaryKey"`
	Clues string `json:"clues"`
}

type SubmitAnswerReq struct {
	Username string `json:"username"`
	ClueID   uint   `json:"clue_id"`
	Answer   string `json:"answer"`
}

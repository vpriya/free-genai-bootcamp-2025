package models

type Group struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	WordCount int64  `json:"word_count,omitempty"`
}

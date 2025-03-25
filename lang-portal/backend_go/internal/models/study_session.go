package models

import "time"

type StudySession struct {
	ID              int64     `json:"id"`
	GroupID         int64     `json:"group_id"`
	StudyActivityID int64     `json:"study_activity_id"`
	Correct         bool      `json:"correct"`
	CreatedAt       time.Time `json:"created_at"`
	GroupName       string    `json:"group_name,omitempty"`
}

type StudyActivity struct {
	ID        int64     `json:"id"`
	GroupID   int64     `json:"group_id"`
	CreatedAt time.Time `json:"created_at"`
}

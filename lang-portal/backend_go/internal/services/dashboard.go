package services

import (
	"database/sql"
)

type DashboardService struct {
	db *sql.DB
}

func NewDashboardService(db *sql.DB) *DashboardService {
	return &DashboardService{db: db}
}

func (s *DashboardService) GetLastStudySession() (map[string]interface{}, error) {
	// Business logic and DB queries here
	return map[string]interface{}{
		"id":               1,
		"group_id":         1,
		"study_activity_id": 1,
		"created_at":       "2024-03-20T15:04:05Z",
		"group_name":       "Basic Greetings",
	}, nil
}

func (s *DashboardService) GetStudyProgress() (map[string]interface{}, error) {
	return map[string]interface{}{
		"total_words_studied": 3,
		"total_words":         124,
	}, nil
}

func (s *DashboardService) GetQuickStats() (map[string]interface{}, error) {
	return map[string]interface{}{
		"success_rate":         0.80,
		"total_study_sessions": 4,
		"total_active_groups":  3,
		"streak_days":         4,
	}, nil
} 
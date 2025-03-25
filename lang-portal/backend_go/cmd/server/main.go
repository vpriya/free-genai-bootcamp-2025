package main

import (
	"database/sql"
	"log"

	"lang-portal/backend_go/internal/handlers"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Initialize DB connection
	db, err := sql.Open("sqlite3", "words.db")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Initialize handler
	h := handlers.NewHandler(db)

	r := gin.Default()

	// Enable CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// API routes
	api := r.Group("/api")
	{
		// Health check
		api.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "ok"})
		})

		// Dashboard routes with actual handlers
		dashboard := api.Group("/dashboard")
		{
			dashboard.GET("/last_study_session", h.GetLastStudySession)
			dashboard.GET("/study_progress", h.GetStudyProgress)
			dashboard.GET("/quick_stats", h.GetQuickStats)
		}

		// Study activities routes
		activities := api.Group("/study_activities")
		{
			activities.GET("/:id", func(c *gin.Context) {})
			activities.GET("/:id/study_sessions", func(c *gin.Context) {})
			activities.POST("", func(c *gin.Context) {})
		}

		// Words routes
		words := api.Group("/words")
		{
			words.GET("", func(c *gin.Context) {})
			words.GET("/:id", func(c *gin.Context) {})
		}

		// Groups routes
		groups := api.Group("/groups")
		{
			groups.GET("", func(c *gin.Context) {})
			groups.GET("/:id", func(c *gin.Context) {})
			groups.GET("/:id/words", func(c *gin.Context) {})
			groups.GET("/:id/study_sessions", func(c *gin.Context) {})
		}

		// Study sessions routes
		sessions := api.Group("/study_sessions")
		{
			sessions.GET("", func(c *gin.Context) {})
			sessions.GET("/:id", func(c *gin.Context) {})
			sessions.GET("/:id/words", func(c *gin.Context) {})
			sessions.POST("/:study_session_id/words/:word_id/review", func(c *gin.Context) {})
		}

		// Reset routes
		api.POST("/reset_history", func(c *gin.Context) {})
		api.POST("/full_reset", func(c *gin.Context) {})
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

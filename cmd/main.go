package main

import (
    "log"
    "sat-siber-recruitment-lms/backend/internal/config"
    "sat-siber-recruitment-lms/backend/internal/handlers"
    "github.com/gin-gonic/gin"
)

func main() {
    db := config.InitDB()
    r := gin.Default()

    // Routes
    r.GET("/api/jobs", handlers.GetJobs)
    r.POST("/api/jobs", handlers.CreateJob)
    r.POST("/api/applications", handlers.SubmitApplication)

    log.Println("Sat Siber Recruitment LMS back-end running on :8080")
    r.Run(":8080")
}

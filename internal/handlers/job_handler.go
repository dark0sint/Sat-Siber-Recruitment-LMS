package handlers

import (
    "net/http"
    "sat-siber-recruitment-lms/backend/internal/models"
    "sat-siber-recruitment-lms/backend/internal/services"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func GetJobs(c *gin.Context) {
    var jobs []models.Job
    db := c.MustGet("db").(*gorm.DB)
    db.Find(&jobs)
    c.JSON(http.StatusOK, jobs)
}

func CreateJob(c *gin.Context) {
    var job models.Job
    if err := c.ShouldBindJSON(&job); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    db := c.MustGet("db").(*gorm.DB)
    db.Create(&job)
    c.JSON(http.StatusOK, job)
}

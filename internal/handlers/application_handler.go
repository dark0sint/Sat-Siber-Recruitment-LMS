package handlers

import (
    "net/http"
    "sat-siber-recruitment-lms/backend/internal/models"
    "sat-siber-recruitment-lms/backend/internal/services"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func SubmitApplication(c *gin.Context) {
    var app models.Application
    if err := c.ShouldBindJSON(&app); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    // Integrate Python for resume processing
    processedResume := services.ProcessResume(app.Resume)
    app.Resume = processedResume
    db := c.MustGet("db").(*gorm.DB)
    db.Create(&app)
    c.JSON(http.StatusOK, app)
}

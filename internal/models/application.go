package models

import "gorm.io/gorm"

type Application struct {
    gorm.Model
    JobID   uint   `json:"job_id"`
    Name    string `json:"name"`
    Email   string `json:"email"`
    Resume  string `json:"resume"` // File path or URL
}

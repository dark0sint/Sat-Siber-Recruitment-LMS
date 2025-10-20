package models

import "gorm.io/gorm"

type Job struct {
    gorm.Model
    Title       string `json:"title"`
    Category    string `json:"category"`
    Location    string `json:"location"`
    Department  string `json:"department"`
    Description string `json:"description"`
}

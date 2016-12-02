package storage

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model

	Name           string
	StartTime      *time.Time
	EndTime        *time.Time
	ExpectDuration time.Duration
	Priority       uint

	Project   Project
	ProjectID int

	Section   Section
	SectinoID int
}

type Project struct {
	gorm.Model

	Name string
}

type Section struct {
	gorm.Model

	StartTime time.Time
}

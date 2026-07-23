package models

import (
	"time"

	"github.com/google/uuid"
)

type JobApplication struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	UserID uuid.UUID `gorm:"column:user_id;type:uuid;not null"`

	CompanyName string `gorm:"column:company_name;not null"`

	JobTitle string `gorm:"column:job_title;not null"`

	JobLink string `gorm:"column:job_link"`

	JobPortal string `gorm:"column:job_portal;not null"`

	Location string `gorm:"column:location"`

	Status string `gorm:"column:status;not null"`

	AppliedDate time.Time `gorm:"column:applied_date"`

	Notes string `gorm:"column:notes"`

	CreatedAt time.Time `gorm:"column:created_at"`

	UpdatedAt time.Time `gorm:"column:updated_at"`
}
package repositories

import (
	"errors"
"strings"
	"github.com/Sagar1329/task-tracker-api/internal/dto"
	"github.com/Sagar1329/task-tracker-api/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)


type JobApplicationRepository struct {
	db *gorm.DB
}

func NewJobApplicationRepository(db *gorm.DB) *JobApplicationRepository{
	return &JobApplicationRepository{
		db: db,
	}
}
//Create job application
func (r *JobApplicationRepository) Create(jobApplication *models.JobApplication) error {
	return r.db.Create(jobApplication).Error

}



// Get user by id

func (r *JobApplicationRepository) GetByIDAndUserID(id uuid.UUID, userID uuid.UUID)(*models.JobApplication,error){
	var jobApplication models.JobApplication

	err := r.db.
		Where("id = ? AND user_id = ?",id,userID).
		First(&jobApplication).Error
	
		if err!=nil{
			if errors.Is(err, gorm.ErrRecordNotFound){
				return nil, nil
			}
			return nil, err
		}

		return &jobApplication, nil

}


func (r *JobApplicationRepository) GetAllByUserID(
	userID uuid.UUID,
	query dto.ListJobApplicationsQuery,
) ([]models.JobApplication, int64, error) {


	var (
		jobApplications []models.JobApplication
		totalItems      int64
	)

    offset := (query.Page - 1) * query.Limit

	db := r.db.Model(&models.JobApplication{}).
		Where("user_id = ?", userID)

		// Search
	if query.Search != "" {
		search := "%" + strings.TrimSpace(query.Search) + "%"

		db = db.Where(
			"(company_name ILIKE ? OR job_title ILIKE ?)",
			search,
			search,
		)
	}

	if query.Status != "" {
	db = db.Where("status = ?", query.Status)
}

// Job Portal Filter
if query.JobPortal != "" {
	db = db.Where("job_portal = ?", query.JobPortal)
}

// Location Filter
if query.Location != "" {
	db = db.Where("location ILIKE ?", "%"+strings.TrimSpace(query.Location)+"%")
}

	if err := db.Count(&totalItems).Error; err != nil {
		return nil, 0, err
	}

	allowedSortFields := map[string]string{
	"company_name": "company_name",
	"job_title":    "job_title",
	"job_portal":   "job_portal",
	"location":     "location",
	"status":       "status",
	"applied_date": "applied_date",
	"created_at":   "created_at",
	"updated_at":   "updated_at",
}

sortField, ok := allowedSortFields[query.Sort]
if !ok {
	sortField = "applied_date"
}

order := strings.ToUpper(query.Order)
if order != "ASC" && order != "DESC" {
	order = "DESC"
}
	if err := db.
		Order(sortField + " " + order).
		Offset(offset).
		Limit(query.Limit).
		Find(&jobApplications).Error; err != nil {
		return nil, 0, err
	}

	return jobApplications, totalItems, nil
}


func (r *JobApplicationRepository) Update(jobApplication *models.JobApplication) error {
	return r.db.Save(jobApplication).Error
}

func (r *JobApplicationRepository) Delete(id uuid.UUID) error {

	return r.db.
		Delete(&models.JobApplication{}, id).
		Error
}
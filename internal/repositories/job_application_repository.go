package repositories

import (
	"github.com/Sagar1329/task-tracker-api/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"errors"
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


func (r *JobApplicationRepository) GetAllByUserID(userId uuid.UUID) ([]models.JobApplication,error){


	var jobApplications []models.JobApplication

	err := r.db.
	Where("user_id = ?",userId).
	Order("applied_date DESC").
	Find(&jobApplications).Error

	if err != nil {
		return nil,err
	}

	return jobApplications, nil
}


func (r *JobApplicationRepository) Update(jobApplication *models.JobApplication) error {
	return r.db.Save(jobApplication).Error
}

func (r *JobApplicationRepository) Delete(id uuid.UUID) error {

	return r.db.
		Delete(&models.JobApplication{}, id).
		Error
}
package services

import (
	"time"
	"fmt"
	 "math"

	"github.com/google/uuid"
appErrors "github.com/Sagar1329/task-tracker-api/internal/errors"
	"github.com/Sagar1329/task-tracker-api/internal/dto"
	"github.com/Sagar1329/task-tracker-api/internal/mapper"
	"github.com/Sagar1329/task-tracker-api/internal/models"
	"github.com/Sagar1329/task-tracker-api/internal/repositories"
)

type JobApplicationService struct {
	jobRepository *repositories.JobApplicationRepository
}

func NewJobApplicationService(
	jobRepository *repositories.JobApplicationRepository,
)*JobApplicationService{
	return &JobApplicationService{
		jobRepository: jobRepository,
	}
}

func (s *JobApplicationService) Create(
	userID uuid.UUID,
	request dto.CreateJobApplicationRequest,
)(*dto.JobApplicationResponse, error){

fmt.Printf("%+v\n", request)
fmt.Println("Applied Date:", request.AppliedDate)
	appliedDate, err := time.Parse("2006-01-02", request.AppliedDate)
if err != nil {
	fmt.Println("jk--->",err)
	return nil, appErrors.ErrInvalidAppliedDate
}
	


	jobApplication := models.JobApplication{
		UserID: userID,
		CompanyName: request.CompanyName,
		JobTitle: request.JobTitle,
		JobLink: request.JobLink,
		JobPortal: request.JobPortal,
		Location: request.Location,
		Status: "Applied",
		AppliedDate: appliedDate,
		Notes: request.Notes,

	}

	if err := s.jobRepository.Create(&jobApplication); err !=nil{
		return nil,err
	}

	response := mapper.ToJobApplicationResponse(&jobApplication)

	return &response,nil
}


func (s *JobApplicationService) GetAll(
	userID uuid.UUID,
	query dto.ListJobApplicationsQuery,
) (*dto.ListJobApplicationsResponse, error) {
	jobApplications, totalItems, err := s.jobRepository.GetAllByUserID(userID, query)

	if err !=nil{
		return nil,err
	}

items := mapper.ToJobApplicationResponses(jobApplications)

	totalPages := int(math.Ceil(float64(totalItems) / float64(query.Limit)))

	return &dto.ListJobApplicationsResponse{
		Items: items,
		Pagination: dto.Pagination{
			Page:       query.Page,
			Limit:      query.Limit,
			TotalItems: int(totalItems),
			TotalPages: totalPages,
		},
	}, nil
}

func (s *JobApplicationService) GetByID(
	userID uuid.UUID,
	jobApplicationID uuid.UUID,
) (*dto.JobApplicationResponse, error){


	jobApplication,err := s.jobRepository.GetByIDAndUserID(
		jobApplicationID,
		userID,
	)

	if err != nil{
		return nil,err
	}

	if jobApplication == nil {
		return nil, appErrors.ErrJobApplicationNotFound
	}

	response := mapper.ToJobApplicationResponse(jobApplication)

	return  &response,nil

}


func (s *JobApplicationService) Update(
	userID uuid.UUID,
	jobApplicationID uuid.UUID,
	request dto.UpdateJobApplicationRequest,
)(*dto.JobApplicationResponse,error){
	
	jobApplication, err := s.jobRepository.GetByIDAndUserID(
		jobApplicationID,
		userID,
	)

	if err != nil{
		return nil,err
	}


	if jobApplication == nil{
		return nil,appErrors.ErrJobApplicationNotFound
	}

	appliedDate, err := time.Parse("2006-01-02", request.AppliedDate)

	if err != nil{
		return nil,appErrors.ErrInvalidAppliedDate
	}

		jobApplication.CompanyName = request.CompanyName
	jobApplication.JobTitle = request.JobTitle
	jobApplication.JobLink = request.JobLink
	jobApplication.JobPortal = request.JobPortal
	jobApplication.Location = request.Location
	jobApplication.Status = request.Status
	jobApplication.AppliedDate = appliedDate
	jobApplication.Notes = request.Notes

	if err := s.jobRepository.Update(jobApplication); err != nil {
		return nil, err
		
	}


	response := mapper.ToJobApplicationResponse(jobApplication)

	return &response, nil
}


func (s *JobApplicationService) Delete(
	userID uuid.UUID,
	jobApplicationID uuid.UUID,
) error{
	jobApplication, err := s.jobRepository.GetByIDAndUserID(
		jobApplicationID,
		userID,
	)

	if err !=nil{
		return err
	}

	if jobApplication == nil{
		return appErrors.ErrJobApplicationNotFound
	}

	if err := s.jobRepository.Delete(jobApplicationID); err !=nil{
		return err
	}

	return nil
}
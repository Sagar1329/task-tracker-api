package mapper
import (
	"time"
"github.com/Sagar1329/task-tracker-api/internal/dto"
	"github.com/Sagar1329/task-tracker-api/internal/models"
)

func ToJobApplicationResponse(
	job *models.JobApplication,
) dto.JobApplicationResponse {

	return dto.JobApplicationResponse{
		ID: job.ID.String(),

		CompanyName: job.CompanyName,

		JobTitle: job.JobTitle,

		JobLink: job.JobLink,

		JobPortal: job.JobPortal,

		Location: job.Location,

		Status: job.Status,

		AppliedDate: job.AppliedDate.Format("2006-01-02"),

		Notes: job.Notes,

		CreatedAt: job.CreatedAt.Format(time.RFC3339),

		UpdatedAt: job.UpdatedAt.Format(time.RFC3339),
	}
}


func ToJobApplicationResponses(
	jobs []models.JobApplication,
) []dto.JobApplicationResponse {

	responses := make([]dto.JobApplicationResponse, 0, len(jobs))

	for _, job := range jobs {

		responses = append(
			responses,
			ToJobApplicationResponse(&job),
		)
	}

	return responses
}
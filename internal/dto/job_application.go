package dto

type CreateJobApplicationRequest struct {
	CompanyName string `json:"company_name" binding:"required"`

	JobTitle string `json:"job_title" binding:"required"`

	JobLink string `json:"job_link"`

	JobPortal string `json:"job_portal" binding:"required"`

	Location string `json:"location"`

	AppliedDate string `json:"applied_date" binding:"required"`

	Notes string `json:"notes"`
}

type UpdateJobApplicationRequest struct {
	CompanyName string `json:"company_name"`
	JobTitle    string `json:"job_title"`
	JobLink     string `json:"job_link"`
	JobPortal   string `json:"job_portal"`
	Location    string `json:"location"`
	Status      string `json:"status"`
	AppliedDate string `json:"applied_date"`
	Notes       string `json:"notes"`
}
type JobApplicationResponse struct {
	ID string `json:"id"`

	CompanyName string `json:"company_name"`

	JobTitle string `json:"job_title"`

	JobLink string `json:"job_link"`

	JobPortal string `json:"job_portal"`

	Location string `json:"location"`

	Status string `json:"status"`

	AppliedDate string `json:"applied_date"`

	Notes string `json:"notes"`

	CreatedAt string `json:"created_at"`

	UpdatedAt string `json:"updated_at"`
}

type ListJobApplicationsResponse struct {
	Applications []JobApplicationResponse `json:"applications"`
	  Total int64 `json:"total"`
}



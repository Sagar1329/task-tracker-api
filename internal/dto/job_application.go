package dto

type CreateJobApplicationRequest struct {
	CompanyName string `json:"company_name" binding:"required,min=2,max=255"`

	JobTitle string `json:"job_title" binding:"required,min=2,max=255"`

	JobLink string `json:"job_link" binding:"omitempty,url,max=2048"`

	JobPortal string `json:"job_portal" binding:"required,oneof=LinkedIn Naukri Indeed Wellfound Other"`

	Location string `json:"location" binding:"omitempty,max=255"`

	AppliedDate string `json:"applied_date" binding:"required,datetime=2006-01-02"`

	Notes string `json:"notes" binding:"omitempty,max=5000"`
}

type UpdateJobApplicationRequest struct {
	CompanyName string `json:"company_name" binding:"required,min=2,max=255"`

	JobTitle string `json:"job_title" binding:"required,min=2,max=255"`

	JobLink string `json:"job_link" binding:"omitempty,url,max=2048"`

	JobPortal string `json:"job_portal" binding:"required,oneof=LinkedIn Naukri Indeed Wellfound Other"`

	Location string `json:"location" binding:"omitempty,max=255"`

	Status string `json:"status" binding:"required,oneof=Applied Interview Rejected Offer"`

	AppliedDate string `json:"applied_date" binding:"required,datetime=2006-01-02"`

	Notes string `json:"notes" binding:"omitempty,max=5000"`
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



type ListJobApplicationsQuery struct {
	Page       int    `form:"page"`
	Limit      int    `form:"limit"`
	Search     string `form:"search"`
	Status     string `form:"status"`
	JobPortal  string `form:"job_portal"`
	Location   string `form:"location"`
	Sort       string `form:"sort"`
	Order      string `form:"order"`
}

func (q *ListJobApplicationsQuery) Normalize() {

	if q.Page < 1 {
		q.Page = 1
	}

	switch {
	case q.Limit < 1:
		q.Limit = 10
	case q.Limit > 100:
		q.Limit = 100
	}

	if q.Sort == "" {
		q.Sort = "applied_date"
	}

	if q.Order == "" {
		q.Order = "desc"
	}
}

type Pagination struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalItems int `json:"total_items"`
	TotalPages int `json:"total_pages"`
}

type ListJobApplicationsResponse struct {
	Items      []JobApplicationResponse `json:"items"`
	Pagination Pagination               `json:"pagination"`
}


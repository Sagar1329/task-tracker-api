package handlers

import (

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/Sagar1329/task-tracker-api/internal/app"
	"github.com/Sagar1329/task-tracker-api/internal/dto"
	"github.com/Sagar1329/task-tracker-api/internal/repositories"
	"github.com/Sagar1329/task-tracker-api/internal/services"
	"github.com/Sagar1329/task-tracker-api/internal/httpresponse"
)

type JobApplicationHandler struct {
	app *app.Application
	jobService *services.JobApplicationService
}

func NewJobApplicationHandler(app *app.Application) *JobApplicationHandler{
	jobRepository := repositories.NewJobApplicationRepository(app.DB)

	jobService := services.NewJobApplicationService(jobRepository)

		return &JobApplicationHandler{
		app:        app,
		jobService: jobService,
	}
}
func getUserID(c *gin.Context) uuid.UUID {
	return c.MustGet("userID").(uuid.UUID)
}
func getJobApplicationID(c *gin.Context) (uuid.UUID, error) {
	return uuid.Parse(c.Param("id"))
}


func (h *JobApplicationHandler) Create(c *gin.Context) {

	var request dto.CreateJobApplicationRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		httpresponse.Error(c, http.StatusBadRequest, httpresponse.FormatValidationError(err))
		return
	}

	userID := getUserID(c)

	jobApplicationResponse, err := h.jobService.Create(
		userID,
		request,
	)

	if err != nil {
		httpresponse.HandleError(c, err)
		return
	}

	httpresponse.Success(
		c,
		http.StatusCreated,
		"Job application created successfully.",
		jobApplicationResponse,
	)
}

func (h *JobApplicationHandler) GetAll(c *gin.Context) {

	userID := getUserID(c)

	jobApplicationsResponse, err := h.jobService.GetAll(userID)

	if err != nil {
		httpresponse.HandleError(c, err)
		return
	}

	httpresponse.Success(
		c,
		http.StatusOK,
		"Job applications fetched successfully.",
		jobApplicationsResponse,
	)
}

func (h *JobApplicationHandler) GetByID(c *gin.Context) {

	jobApplicationID, err := getJobApplicationID(c)
	if err != nil {
		httpresponse.Error(c, http.StatusBadRequest, "invalid job application id")
		return
	}

	userID := getUserID(c)

	jobApplicationResponse, err := h.jobService.GetByID(
		userID,
		jobApplicationID,
	)

	if err != nil {
		httpresponse.HandleError(c, err)
		return
	}

	httpresponse.Success(
		c,
		http.StatusOK,
		"Job application fetched successfully.",
		jobApplicationResponse,
	)
}



func (h *JobApplicationHandler) Update(c *gin.Context) {

	var request dto.UpdateJobApplicationRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		httpresponse.Error(c, http.StatusBadRequest, httpresponse.FormatValidationError(err))
		return
	}

	jobApplicationID, err := getJobApplicationID(c)
	if err != nil {
		httpresponse.Error(c, http.StatusBadRequest, "invalid job application id")
		return
	}

	userID := getUserID(c)

	jobApplicationResponse, err := h.jobService.Update(
		userID,
		jobApplicationID,
		request,
	)

	if err != nil {
		httpresponse.HandleError(c, err)
		return
	}

	httpresponse.Success(
		c,
		http.StatusOK,
		"Job application updated successfully.",
		jobApplicationResponse,
	)
}


func (h *JobApplicationHandler) Delete(c *gin.Context) {

	jobApplicationID, err := getJobApplicationID(c)
	if err != nil {
		httpresponse.Error(c, http.StatusBadRequest, "invalid job application id")
		return
	}

	userID := getUserID(c)

	err = h.jobService.Delete(userID, jobApplicationID)
	if err != nil {
		httpresponse.HandleError(c, err)
		return
	}

	httpresponse.Success(
		c,
		http.StatusOK,
		"Job application deleted successfully.",
		nil,
	)
}
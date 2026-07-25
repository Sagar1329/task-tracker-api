package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/Sagar1329/task-tracker-api/internal/app"
	appErrors "github.com/Sagar1329/task-tracker-api/internal/errors"
	"github.com/Sagar1329/task-tracker-api/internal/dto"
	"github.com/Sagar1329/task-tracker-api/internal/repositories"
	"github.com/Sagar1329/task-tracker-api/internal/services"
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

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	userID := getUserID(c)

	response, err := h.jobService.Create(
		userID,
		request,
	)

	if err != nil {

		if errors.Is(err, appErrors.ErrInvalidAppliedDate) {

			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})

		return
	}

	c.JSON(http.StatusCreated, response)
}


func (h *JobApplicationHandler) GetAll(c *gin.Context) {

	userID := getUserID(c)



	response, err := h.jobService.GetAll(userID)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})

		return
	}

	c.JSON(http.StatusOK, response)
}


func (h *JobApplicationHandler) GetByID(c *gin.Context) {

jobApplicationID, err := getJobApplicationID(c)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid job application id",
		})

		return
	}

	userID := getUserID(c)

	response, err := h.jobService.GetByID(
		userID,
		jobApplicationID,
	)

	if err != nil {

		if errors.Is(err, appErrors.ErrJobApplicationNotFound) {

			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})

		return
	}

	c.JSON(http.StatusOK, response)
}


func (h *JobApplicationHandler) Update(c *gin.Context) {

	var request dto.UpdateJobApplicationRequest

	if err := c.ShouldBindJSON(&request); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	jobApplicationID, err := getJobApplicationID(c)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid job application id",
		})

		return
	}

	userID := getUserID(c)

	response, err := h.jobService.Update(
		userID,
		jobApplicationID,
		request,
	)

	if err != nil {

		switch {

		case errors.Is(err, appErrors.ErrJobApplicationNotFound):

			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})

			return

		case errors.Is(err, appErrors.ErrInvalidAppliedDate):

			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return

		default:

			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal Server Error",
			})

			return
		}
	}

	c.JSON(http.StatusOK, response)
}


func (h *JobApplicationHandler) Delete(c *gin.Context) {

	jobApplicationID, err := getJobApplicationID(c)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid job application id",
		})

		return
	}

	userID := getUserID(c)

	err = h.jobService.Delete(
		userID,
		jobApplicationID,
	)

	if err != nil {

		if errors.Is(err, appErrors.ErrJobApplicationNotFound) {

			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})

		return
	}

	c.Status(http.StatusNoContent)
}
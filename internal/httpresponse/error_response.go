package httpresponse

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	appErrors "github.com/Sagar1329/task-tracker-api/internal/errors"
)

func HandleError(c *gin.Context, err error) {

	switch {

	case errors.Is(err, appErrors.ErrInvalidCredentials):
		Error(c, http.StatusUnauthorized, err.Error())

	case errors.Is(err, appErrors.ErrEmailAlreadyExists):
		Error(c, http.StatusConflict, err.Error())

	case errors.Is(err, appErrors.ErrJobApplicationNotFound):
		Error(c, http.StatusNotFound, err.Error())

	case errors.Is(err, appErrors.ErrInvalidAppliedDate):
		Error(c, http.StatusBadRequest, err.Error())

	default:
		Error(c, http.StatusInternalServerError, "Internal Server Error")
	}
}
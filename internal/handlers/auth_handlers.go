package handlers

import (
	"errors"
	"net/http"

	appErrors "github.com/Sagar1329/task-tracker-api/internal/errors"

	"github.com/Sagar1329/task-tracker-api/internal/app"
	"github.com/Sagar1329/task-tracker-api/internal/dto"
	"github.com/Sagar1329/task-tracker-api/internal/repositories"
	"github.com/Sagar1329/task-tracker-api/internal/services"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	app         *app.Application
	authService *services.AuthService
}


func NewAuthHandler(app *app.Application) *AuthHandler{
	userRepository := repositories.NewUserRepository(app.DB)

	authService := services.NewAuthService(
		userRepository,
		app.Config.JWT.Secret,
	)

	return &AuthHandler{
		app: app,
		authService: authService,
	}
}

func (h *AuthHandler) Signup(c *gin.Context){
	var request dto.SignupRequest

	if err := c.ShouldBindJSON(&request); err !=nil{

		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
		return
	}

	err := h.authService.Signup(request)

	if err !=nil{
		if errors.Is(err, appErrors.ErrUserAlreadyExists){

			c.JSON(http.StatusConflict, gin.H{
				"error": err.Error(),
			})

			return
		}
	
			c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})

		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
	})
}


func (h *AuthHandler) Login(c *gin.Context) {

	var request dto.LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	response, err := h.authService.Login(request)

	if err != nil {

		if errors.Is(err, appErrors.ErrInvalidCredentials) {

			c.JSON(http.StatusUnauthorized, gin.H{
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
package services

import (
	
appErrors "github.com/Sagar1329/task-tracker-api/internal/errors"
	"github.com/Sagar1329/task-tracker-api/internal/auth"
	"github.com/Sagar1329/task-tracker-api/internal/dto"
	"github.com/Sagar1329/task-tracker-api/internal/models"
	"github.com/Sagar1329/task-tracker-api/internal/repositories"
)

type AuthService struct {
    userRepository *repositories.UserRepository
    jwtSecret      string
}

func NewAuthService(
	userRepository *repositories.UserRepository,
	jwtSecret string,
) *AuthService {

	return &AuthService{
		userRepository: userRepository,
		jwtSecret:      jwtSecret,
	}
}



func (s *AuthService) Signup(request dto.SignupRequest) error {


	existingUser,err := s.userRepository.GetByEmail(request.Email)

	if err != nil {
		return err
	}

	if existingUser != nil{
		return appErrors.ErrUserAlreadyExists
	}

	hashedPassword, err := auth.HashPassword(request.Password)

	if err !=nil{
		return err
	}

	user := models.User{
		Name: request.Name,
		Email: request.Email,
		PasswordHash: hashedPassword,
	}

	if err := s.userRepository.CreateUser(&user); err !=nil{
		return nil
	}

	return nil
}


func (s *AuthService) Login(request dto.LoginRequest) (*dto.LoginResponse, error) {

	// Find user by email
	user, err := s.userRepository.GetByEmail(request.Email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, appErrors.ErrInvalidCredentials
	}

	// Verify password
	if !auth.VerifyPassword(request.Password, user.PasswordHash) {
		return nil, appErrors.ErrInvalidCredentials
	}

	// Generate JWT
	token, err := auth.GenerateToken(
		user.ID,
		user.Email,
		s.jwtSecret,
	)

	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		AccessToken: token,
		ExpiresIn:   int64(auth.TokenExpiry.Seconds()),
	}, nil
}
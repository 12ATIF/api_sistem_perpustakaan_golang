package handler

import (
	"coba_dulu/internal/entity"
	"coba_dulu/internal/usecase"
	"coba_dulu/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
	jwtService  pkg.JWTService
}

func NewUserHandler(userUsecase usecase.UserUsecase, jwtService pkg.JWTService) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
		jwtService:  jwtService,
	}
}

type UserCreateDTO struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

type UserLoginDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (h *UserHandler) Register(ctx *gin.Context) {
	var userDTO UserCreateDTO
	if err := ctx.ShouldBindJSON(&userDTO); err != nil {
		ValidationResponse(ctx, "Invalid request", SplitError(err))
		return
	}

	user := entity.User{
		Name:     userDTO.Name,
		Email:    userDTO.Email,
		Password: userDTO.Password,
		Role:     userDTO.Role,
	}

	createdUser, err := h.userUsecase.CreateUser(user)
	if err != nil {
		ErrorResponse(ctx, "Failed to create user", http.StatusInternalServerError, err.Error())
		return
	}
	SuccessResponse(ctx, "User created successfully", createdUser)
}

func (h *UserHandler) Login(ctx *gin.Context) {
	var loginDTO UserLoginDTO
	if err := ctx.ShouldBindJSON(&loginDTO); err != nil {
		ValidationResponse(ctx, "Invalid request", SplitError(err))
		return
	}

	authedUser, err := h.userUsecase.Login(loginDTO.Email, loginDTO.Password)
	if err != nil {
		ErrorResponse(ctx, "Invalid credentials", http.StatusUnauthorized, "Invalid email or password")
		return
	}
	
	user := authedUser.(entity.User)
	token, err := h.jwtService.GenerateToken(user.ID, user.Role)
	if err != nil {
		ErrorResponse(ctx, "Failed to generate token", http.StatusInternalServerError, err.Error())
		return
	}

	response := gin.H{
		"token": token,
	}
	SuccessResponse(ctx, "Login successful", response)
}

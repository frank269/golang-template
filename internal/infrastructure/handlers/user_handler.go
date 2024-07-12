package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/mpt-tiendc/golang-template/internal/config"
	"github.com/mpt-tiendc/golang-template/internal/domain/dtos"
	"github.com/mpt-tiendc/golang-template/internal/domain/entities"
	"github.com/mpt-tiendc/golang-template/internal/domain/services"
	"github.com/mpt-tiendc/golang-template/internal/infrastructure/database"
	repositoriesimpl "github.com/mpt-tiendc/golang-template/internal/infrastructure/repositories_impl"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(appConfig *config.AppConfig) (*UserHandler, error) {
	db, err := database.NewDatabase(appConfig.DatabaseConfig)
	if err != nil {
		return nil, err
	}

	return &UserHandler{
		userService: services.NewUserService(repositoriesimpl.NewUserRepositoryImpl(db)),
	}, nil
}

func (h *UserHandler) GetUserById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := h.userService.GetUserById(chi.URLParam(r, "id"))
		if err != nil {
			render.JSON(w, r, dtos.BaseResponse{
				Success: false,
				Message: err.Error(),
			})
			return
		}
		render.JSON(w, r, dtos.BaseResponse{
			Success: true,
			Message: "get user success",
			Data:    user,
		})
	}
}

func (h *UserHandler) CreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request dtos.CreateUserRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			render.JSON(w, r, dtos.BaseResponse{
				Success: false,
				Message: err.Error(),
			})
			return
		}
		user := &entities.User{
			Username: request.Username,
			Password: &request.Password, // generate pass is here
		}
		err = h.userService.Create(user)
		if err != nil {
			render.JSON(w, r, dtos.BaseResponse{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		render.JSON(w, r, dtos.BaseResponse{
			Success: true,
			Message: "create user success",
			Data:    user,
		})
	}
}

func (h *UserHandler) UpdateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := h.userService.GetUserById(chi.URLParam(r, "id"))
		if err != nil {
			render.JSON(w, r, dtos.BaseResponse{
				Success: false,
				Message: err.Error(),
			})
			return
		}
		var request dtos.UpdateUserRequest
		err = json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			render.JSON(w, r, dtos.BaseResponse{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		user.Password = &request.Password // generate pass is here

		err = h.userService.Update(user)
		if err != nil {
			render.JSON(w, r, dtos.BaseResponse{
				Success: false,
				Message: err.Error(),
			})
			return
		}
		render.JSON(w, r, dtos.BaseResponse{
			Success: true,
			Message: "update user success",
			Data:    user,
		})
	}
}

func (h *UserHandler) DeleteUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := h.userService.GetUserById(chi.URLParam(r, "id"))
		if err != nil {
			render.JSON(w, r, dtos.BaseResponse{
				Success: false,
				Message: err.Error(),
			})
			return
		}
		err = h.userService.Delete(user.Id)
		if err != nil {
			render.JSON(w, r, dtos.BaseResponse{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		render.JSON(w, r, dtos.BaseResponse{
			Success: true,
			Message: "delete user success",
			Data:    user,
		})
	}
}

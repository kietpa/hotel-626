package handlers

import (
	"hotel/service"
	"hotel/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

// @Summary      Register
// @Description  Register a new user
// @Tags         User
// @Accept       json
// @Produce      json
// @Param		 data body model.User true "The input user struct"
// @Success      201  {object}  model.RegisterResponse
// @Failure      400  {object}  []utils.ApiError
// @Failure      500  {object}  utils.ErrorMssg
// @Router       /users/register [Post]
func (h *Handler) RegisterHandler(c echo.Context) error {
	var input service.RegisterInput

	err := c.Bind(&input)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrFailedBind)
	}

	userInfo, err := h.Service.Register(input)
	if err != nil {
		apiErr := utils.FromError(err)
		return echo.NewHTTPError(apiErr.Status, apiErr.Message)
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "register success",
		"user":    userInfo,
	})
}

// @Summary      Login
// @Description  Login with email and password
// @Tags         User
// @Accept       json
// @Produce      json
// @Param		 data body model.UserLogin true "The input user struct"
// @Success      200  {object}  model.LoginResponse
// @Failure      400  {object}  []utils.ApiError
// @Failure      401  {object}  utils.ErrorMssg
// @Failure      500  {object}  utils.ErrorMssg
// @Router       /api/users/login [Post]
func (h *Handler) LoginHandler(c echo.Context) error {
	var input service.LoginInput

	err := c.Bind(&input)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrFailedBind)
	}

	token, err := h.Service.Login(input)
	if err != nil {
		apiErr := utils.FromError(err)
		return echo.NewHTTPError(apiErr.Status, apiErr.Message)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "login success",
		"token":   token,
	})
}

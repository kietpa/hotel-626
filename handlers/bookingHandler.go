package handlers

import (
	"hotel/helpers"
	"hotel/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) ShowBookingHandler(c echo.Context) error {
	// get logged in user id
	userID, err := helpers.GetUserId(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, utils.ErrInternalFailure)
	}

	// get bookings
	bookings, err := h.Service.ShowUserBookings(userID)
	if err != nil {
		apiErr := utils.FromError(err)
		return echo.NewHTTPError(apiErr.Status, apiErr.Message)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"bookings": bookings,
	})
}
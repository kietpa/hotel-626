package handlers

import (
	"hotel/helpers"
	"hotel/service"
	"hotel/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) RoomTypeHandler(c echo.Context) error {
	roomTypes, err := h.Service.GetRoomTypes()
	if err != nil {
		apiErr := utils.FromError(err)
		return echo.NewHTTPError(apiErr.Status, apiErr.Message)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"room_types": roomTypes,
	})
}

func (h *Handler) AvailableRoomHandler(c echo.Context) error {
	id := c.Param("id")

	rooms, err := h.Service.GetAvailableRooms(id)
	if err != nil {
		apiErr := utils.FromError(err)
		return echo.NewHTTPError(apiErr.Status, apiErr.Message)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"rooms": rooms,
	})
}

func (h *Handler) RoomBookingHandler(c echo.Context) error {
	// bind user input
	var input service.BookingInput
	err := c.Bind(&input)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrFailedBind)
	}

	// get logged in user id
	userID, err := helpers.GetUserId(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, utils.ErrInternalFailure)
	}
	input.User_id = userID

	// book room
	Booking, err := h.Service.BookRoom(input)
	if err != nil {
		apiErr := utils.FromError(err)
		return echo.NewHTTPError(apiErr.Status, apiErr.Message)
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "booking created",
		"booking": Booking,
	})
}
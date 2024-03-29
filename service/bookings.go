package service

import (
	"fmt"
	"hotel/helpers"
	"hotel/model"
	"hotel/utils"
)

func (s *Service) ShowUserBookings(userID int) ([]model.Booking, error) {
	bookings := []model.Booking{}

	// tambah available rooms count?
	err := s.DB.Model(model.Booking{}).Where("user_id = ?", userID).Find(&bookings).Error
	if err != nil {
		return nil, utils.NewError(utils.ErrInternalFailure, err)
	}

	return bookings, nil
}

func (s *Service) BookRoom(input BookingInput) (model.Booking, error) {
	// get room info
	room := model.Room{}
	err := s.DB.Model(model.Room{}).Preload("Room_type").First(&room, input.Room_id).Error
	if err != nil {
		return model.Booking{}, utils.NewError(utils.ErrInternalFailure, err)
	}
	if room.Room_id == 0 {
		return model.Booking{}, utils.NewError(utils.ErrNotFound, fmt.Errorf("room not found"))
	}

	// calculate total price
	days, err := helpers.DatesBetween(input.Checkin_date, input.Checkout_date)
	if err != nil {
		return model.Booking{}, utils.NewError(utils.ErrBadRequest, err) // bad request because date is not correct format
	}

	totalPrice := days * room.Room_type.Price_per_night

	// insert booking
	booking := model.Booking{
		User_id:       input.User_id,
		Room_id:       input.Room_id,
		Checkin_date:  input.Checkin_date,
		Checkout_date: input.Checkout_date,
		Total_price:   totalPrice,
		Paid:          false,
	}

	err = s.DB.Omit("booking_id").Create(&booking).Error
	if err != nil {
		return model.Booking{}, utils.NewError(utils.ErrInternalFailure, err)
	}

	// update room
	err = s.DB.Model(&room).Update("available", false).Error
	if err != nil {
		return model.Booking{}, utils.NewError(utils.ErrInternalFailure, err)
	}

	return booking, nil
}

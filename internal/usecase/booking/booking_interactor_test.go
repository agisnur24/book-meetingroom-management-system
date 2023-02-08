package booking_test

import (
	"book-meeting-hotel/domain/entity"
	mock3 "book-meeting-hotel/internal/repository/mock"
	booking2 "book-meeting-hotel/internal/usecase/booking"
	"github.com/stretchr/testify/assert"
	mock2 "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestItShouldBeReturnNewBooking(t *testing.T) {
	meetingRoomId := 1
	employeeId := 1
	picContactInformation := "08121211221"
	startDatetime := time.Now().Add(10 * 24 * time.Hour)
	endDatetime := startDatetime.Add(4 * time.Hour)

	repo := new(mock3.BookingRepositoryMock)
	repo.On("GetByDateAndMeetingRoom", meetingRoomId, startDatetime, endDatetime).
		Return((*entity.Booking)(nil), nil)
	repo.On("Save", mock2.AnythingOfType("*entity.Booking")).Return(nil)
	meetingRoomRepo := new(mock3.MeetingRoomRepositoryMock)
	meetingRoomRepo.On("GetById", meetingRoomId).Return(&entity.MeetingRoom{
		Id:          1,
		Name:        "Cendrawasih",
		Capacity:    100,
		RatePerDay:  1000000,
		RatePerHour: 20000,
	}, nil)
	interactor := booking2.NewBookingInteract(repo, meetingRoomRepo)

	newbooking, err := interactor.BookMeetingRoom(employeeId, meetingRoomId,
		startDatetime, endDatetime, picContactInformation)

	require.NoError(t, err)
	require.NotNil(t, newbooking)
	assert.Equal(t, 80000, newbooking.GetGrandTotal())
}

func TestItShouldBeReturnErrorWhenMeetingRoomNotFound(t *testing.T) {
	meetingRoomId := 1
	employeeId := 1
	picContactInformation := "08121211221"
	startDatetime := time.Now().Add(10 * 24 * time.Hour)
	endDatetime := startDatetime.Add(4 * time.Hour)

	repo := new(mock3.BookingRepositoryMock)
	meetingRoomRepo := new(mock3.MeetingRoomRepositoryMock)
	meetingRoomRepo.On("GetById", meetingRoomId).Return((*entity.MeetingRoom)(nil), nil)
	interactor := booking2.NewBookingInteract(repo, meetingRoomRepo)
	newbooking, err := interactor.BookMeetingRoom(employeeId, meetingRoomId,
		startDatetime, endDatetime, picContactInformation)

	require.Nil(t, newbooking)
	require.Error(t, err)
}

func TestItShouldBeReturnErrorWhenBookingAlreadyExists(t *testing.T) {
	meetingRoomId := 1
	employeeId := 1
	picContactInformation := "08121211221"
	startDatetime := time.Now().Add(10 * 24 * time.Hour)
	endDatetime := startDatetime.Add(4 * time.Hour)
	meetingRoom := &entity.MeetingRoom{
		Id:          1,
		Name:        "Cendrawasih",
		Capacity:    100,
		RatePerDay:  1000000,
		RatePerHour: 20000,
	}
	repo := new(mock3.BookingRepositoryMock)
	existBooking, _ := entity.NewBooking(employeeId, *meetingRoom, "invoice.jpg", picContactInformation,
		startDatetime, startDatetime.Add(2*time.Hour))
	repo.On("GetByDateAndMeetingRoom", meetingRoomId, startDatetime, endDatetime).
		Return(existBooking, nil)
	meetingRoomRepo := new(mock3.MeetingRoomRepositoryMock)
	meetingRoomRepo.On("GetById", meetingRoomId).Return(meetingRoom, nil)
	interactor := booking2.NewBookingInteract(repo, meetingRoomRepo)
	newbooking, err := interactor.BookMeetingRoom(employeeId, meetingRoomId,
		startDatetime, endDatetime, picContactInformation)

	require.Nil(t, newbooking)
	require.Error(t, err)
	require.ErrorIs(t, err, booking2.ErrMeetingRoomAlreadyBooked)
}

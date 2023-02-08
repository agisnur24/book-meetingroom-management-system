package mock

import (
	"book-meeting-hotel/domain/entity"
	"github.com/stretchr/testify/mock"
	"time"
)

type BookingRepositoryMock struct {
	mock.Mock
}

func (r *BookingRepositoryMock) Save(booking *entity.Booking) error {
	args := r.Called(booking)
	return args.Error(0)
}

func (r *BookingRepositoryMock) GetByDateAndMeetingRoom(meetingRoomId int, startDatetime time.Time,
	endDatetime time.Time) (*entity.Booking, error) {
	args := r.Called(meetingRoomId, startDatetime, endDatetime)
	return args.Get(0).(*entity.Booking), args.Error(1)
}

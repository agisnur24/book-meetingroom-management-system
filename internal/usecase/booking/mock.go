package booking

import (
	"book-meeting-hotel/domain/entity"
	"github.com/stretchr/testify/mock"
	"time"
)

type UseCaseBookingMock struct {
	mock.Mock
}

func (u *UseCaseBookingMock) BookMeetingRoom(meetingRoomId int, employeeId int, startDatetime time.Time,
	endDatetime time.Time, picContactInformation string) (*entity.Booking, error) {
	args := u.Called(meetingRoomId, employeeId, startDatetime, endDatetime, picContactInformation)

	return args.Get(0).(*entity.Booking), args.Error(1)
}

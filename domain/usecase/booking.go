package usecase

import (
	"book-meeting-hotel/domain/entity"
	"time"
)

type BookingUseCase interface {
	BookMeetingRoom(meetingRoomId int, employeeId int, startDatetime time.Time,
		endDatetime time.Time, picContactInformation string) (*entity.Booking, error)
}

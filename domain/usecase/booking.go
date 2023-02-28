package usecase

import (
	"github.com/agisnur24/book-meetingroom-management-system.git/domain/entity"
	"time"
)

type BookingUseCase interface {
	BookMeetingRoom(meetingRoomId int, employeeId int, startDatetime time.Time,
		endDatetime time.Time, picContactInformation string) (*entity.Booking, error)
}

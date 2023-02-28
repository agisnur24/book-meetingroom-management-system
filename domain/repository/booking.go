package repository

import (
	"github.com/agisnur24/book-meetingroom-management-system.git/domain/entity"
	"time"
)

type BookingRepository interface {
	Save(booking *entity.Booking) error
	GetByDateAndMeetingRoom(meetingRoomId int, startDatetime time.Time, endDatetime time.Time) (*entity.Booking, error)
}

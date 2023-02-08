package repository

import (
	"book-meeting-hotel/domain/entity"
	"time"
)

type BookingRepository interface {
	Save(booking *entity.Booking) error
	GetByDateAndMeetingRoom(meetingRoomId int, startDatetime time.Time, endDatetime time.Time) (*entity.Booking, error)
}

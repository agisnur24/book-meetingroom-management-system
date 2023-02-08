package repository

import "book-meeting-hotel/domain/entity"

type MeetingRoomRepository interface {
	GetById(id int) (*entity.MeetingRoom, error)
}

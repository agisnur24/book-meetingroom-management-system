package repository

import "github.com/agisnur24/book-meetingroom-management-system.git/domain/entity"

type MeetingRoomRepository interface {
	GetById(id int) (*entity.MeetingRoom, error)
}

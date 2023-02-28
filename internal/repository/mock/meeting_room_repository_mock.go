package mock

import (
	"github.com/agisnur24/book-meetingroom-management-system.git/domain/entity"
	"github.com/stretchr/testify/mock"
)

type MeetingRoomRepositoryMock struct {
	mock.Mock
}

func (r *MeetingRoomRepositoryMock) GetById(id int) (*entity.MeetingRoom, error) {
	args := r.Called(id)
	return args.Get(0).(*entity.MeetingRoom), args.Error(1)
}

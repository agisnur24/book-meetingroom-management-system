package booking

import (
	"errors"
	"github.com/agisnur24/book-meetingroom-management-system.git/domain/entity"
	"github.com/agisnur24/book-meetingroom-management-system.git/domain/repository"
	"time"
)

type UseCaseBookingInteract struct {
	bookingRepository     repository.BookingRepository
	meetingRoomRepository repository.MeetingRoomRepository
}

func NewBookingInteract(repo repository.BookingRepository,
	meetingRoomRepo repository.MeetingRoomRepository) UseCaseBookingInteract {
	return UseCaseBookingInteract{bookingRepository: repo, meetingRoomRepository: meetingRoomRepo}
}

func (i UseCaseBookingInteract) BookMeetingRoom(meetingRoomId int, employeeId int, startDatetime time.Time,
	endDatetime time.Time, picContactInformation string) (*entity.Booking, error) {
	meetingRoom, errGetMeetingRoom := i.meetingRoomRepository.GetById(meetingRoomId)

	if errGetMeetingRoom != nil {
		return nil, errGetMeetingRoom
	}

	if meetingRoom == nil {
		return nil, errors.New("meeting room not found")
	}

	existBooking, errExistBooking := i.bookingRepository.GetByDateAndMeetingRoom(meetingRoomId, startDatetime,
		endDatetime)
	if errExistBooking != nil {
		return nil, errExistBooking
	}
	if existBooking != nil {
		return nil, ErrMeetingRoomAlreadyBooked
	}

	newBooking, err := entity.NewBooking(employeeId, *meetingRoom, "invoice.jpg",
		picContactInformation, startDatetime, endDatetime)
	if err != nil {
		return nil, err
	}

	err = i.bookingRepository.Save(newBooking)
	if err != nil {
		return nil, err
	}

	return newBooking, err
}

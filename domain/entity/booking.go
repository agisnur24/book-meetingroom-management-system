package entity

import (
	"errors"
	"math"
	"time"
)

type Booking struct {
	Id            int
	EmployeeId    int
	MeetingRoom   MeetingRoom
	PhotoInvoice  string
	ContactPIC    string
	StartDatetime time.Time
	EndDatetime   time.Time
	Status        string
	RatePerDay    int
	RatePerHour   int
	discount      int
	BookDate      time.Time
}

const MAX_CANCELHOUR = 24 * 7

func NewBooking(employeId int, meetingRoom MeetingRoom, photoInvoice string, contactPIC string,
	startDatetime time.Time, endDatetime time.Time) (*Booking, error) {

	return &Booking{
		EmployeeId:    employeId,
		MeetingRoom:   meetingRoom,
		PhotoInvoice:  photoInvoice,
		ContactPIC:    contactPIC,
		StartDatetime: startDatetime,
		EndDatetime:   endDatetime,
		Status:        "BOOK",
		RatePerDay:    meetingRoom.RatePerDay,
		RatePerHour:   meetingRoom.RatePerHour,
		discount:      0,
		BookDate:      time.Now(),
	}, nil
}

func RebuildBooking(id int, employeId int, meetingRoomId int, photoInvoice string, contactPIC string,
	startDatetime time.Time, endDatetime time.Time, status string, ratePerday int, ratePerHour int, discount int,
	bookDate time.Time) *Booking {
	return &Booking{
		Id:            id,
		EmployeeId:    employeId,
		MeetingRoom:   MeetingRoom{},
		PhotoInvoice:  photoInvoice,
		ContactPIC:    contactPIC,
		StartDatetime: startDatetime,
		EndDatetime:   endDatetime,
		Status:        status,
		RatePerDay:    ratePerday,
		RatePerHour:   ratePerHour,
		discount:      discount,
		BookDate:      bookDate,
	}
}

func (b *Booking) GetStatus() string {
	return b.Status
}

func (b *Booking) GetDurationInHour() int {
	durationInhour := b.EndDatetime.Sub(b.StartDatetime).Hours()
	return int(durationInhour)
}

func (b *Booking) GetRatePerDay() int {
	return b.RatePerDay
}

func (b *Booking) GetRatePerHour() int {
	return b.RatePerHour
}

func (b *Booking) GetDiscount() int {
	return b.discount
}

func (b *Booking) GetTotal() int {
	return b.GetDurationInHour() * b.RatePerHour
}

func (b *Booking) GetGrandTotal() int {
	return b.GetTotal() - b.discount
}

func (b *Booking) GetBookDate() string {
	return b.BookDate.Format("2006-01-02")
}

func (b *Booking) Cancel() error {
	if math.Ceil(b.EndDatetime.Sub(time.Now()).Hours()) < (MAX_CANCELHOUR) {
		return errors.New("cannot cancel booking before 1 week")
	}
	b.Status = "CANCEL"
	return nil
}

func (b *Booking) AddDiscount(discount int) error {
	if discount >= b.GetGrandTotal() {
		return errors.New("cannot add discount more than grand total")
	}
	b.discount = discount
	return nil
}

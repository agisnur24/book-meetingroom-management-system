package mapper

import (
	"book-meeting-hotel/domain/entity"
	"book-meeting-hotel/internal/repository/mysql/model"
)

type BookingMapper struct {
}

func (BookingMapper) FromDomain(booking *entity.Booking) *model.BookingMysqlModel {
	return &model.BookingMysqlModel{
		ID:            int32(booking.Id),
		EmployeeId:    int32(booking.EmployeeId),
		MeetingRoomId: int32(booking.MeetingRoom.Id),
		PhotoInvoice:  booking.PhotoInvoice,
		ContactPic:    booking.ContactPIC,
		StartDateTime: booking.StartDatetime,
		EndDateTime:   booking.EndDatetime,
		Status:        booking.Status,
		RatePerDay:    int32(booking.RatePerDay),
		RatePerHour:   int32(booking.RatePerHour),
		Discount:      int32(booking.GetDiscount()),
		BookDate:      booking.BookDate,
	}
}

func (BookingMapper) ToDomain(mysqlModel *model.BookingMysqlModel) *entity.Booking {
	booking := entity.RebuildBooking(int(mysqlModel.ID), int(mysqlModel.EmployeeId), int(mysqlModel.MeetingRoomId),
		mysqlModel.PhotoInvoice, mysqlModel.ContactPic, mysqlModel.StartDateTime, mysqlModel.EndDateTime,
		mysqlModel.Status, int(mysqlModel.RatePerDay), int(mysqlModel.RatePerHour), int(mysqlModel.Discount),
		mysqlModel.BookDate)
	return booking
}

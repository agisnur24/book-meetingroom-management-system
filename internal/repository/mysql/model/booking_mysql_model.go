package model

import "time"

type BookingMysqlModel struct {
	ID            int32     `dbq:"id"`
	EmployeeId    int32     `dbq:"employee_id"`
	MeetingRoomId int32     `dbq:"meeting_room_id"`
	PhotoInvoice  string    `dbq:"photo_invoice"`
	ContactPic    string    `dbq:"contact_pic"`
	StartDateTime time.Time `dbq:"start_datetime"`
	EndDateTime   time.Time `dbq:"end_datetime"`
	Status        string    `dbq:"status"`
	RatePerDay    int32     `dbq:"rate_per_day"`
	RatePerHour   int32     `dbq:"rate_per_hour"`
	Discount      int32     `dbq:"discount"`
	BookDate      time.Time `dbq:"book_date"`
}

func (b BookingMysqlModel) GetTableName() string {
	return "booking"
}

func FieldTable() []string {
	return []string{
		"id",
		"employee_id",
		"meeting_room_id",
		"photo_invoice",
		"contact_pic",
		"start_datetime",
		"end_datetime",
		"status",
		"rate_per_day",
		"rate_per_hour",
		"book_date",
	}
}

package entity_test

import (
	"book-meeting-hotel/domain/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func Test_ItShouldBeSuccessNewBooking(t *testing.T) {
	employeeId := 1
	meetingRoom := entity.MeetingRoom{
		Id:          1,
		Name:        "Room A",
		Capacity:    100,
		RatePerDay:  1000000,
		RatePerHour: 200000,
	}
	photoInvoice := "invoice.jpg"
	startDate := time.Now().Add(24 * time.Hour)
	endDate := startDate.Add(3 * time.Hour)
	contactPIC := "081313131313"
	booking, _ := entity.NewBooking(employeeId, meetingRoom, photoInvoice, contactPIC, startDate, endDate)
	assert.Equal(t, "BOOK", booking.GetStatus())
	assert.Equal(t, 3, booking.GetDurationInHour())
	assert.Equal(t, 1000000, booking.GetRatePerDay())
	assert.Equal(t, 200000, booking.GetRatePerHour())
	assert.Equal(t, 0, booking.GetDiscount())
	assert.Equal(t, 600000, booking.GetTotal())
	assert.Equal(t, 600000, booking.GetGrandTotal())
	expectedBookDate := time.Now().Format("2006-01-02")
	assert.Equal(t, expectedBookDate, booking.GetBookDate())
}

func Test_ItShouldSuccessCancelBooking(t *testing.T) {
	employeeId := 1
	meetingRoom := entity.MeetingRoom{
		Id:          1,
		Name:        "Room A",
		Capacity:    100,
		RatePerDay:  1000000,
		RatePerHour: 200000,
	}
	photoInvoice := "invoice.jpg"
	startDate := time.Now().Add(8 * 24 * time.Hour)
	endDate := startDate.Add(3 * time.Hour)
	contactPIC := "081313131313"
	booking, _ := entity.NewBooking(employeeId, meetingRoom, photoInvoice, contactPIC, startDate, endDate)
	err := booking.Cancel()

	assert.NoError(t, err)
	assert.Equal(t, "CANCEL", booking.GetStatus())
}

func Test_ItShouldFailedToCancelBeforeOneWeek(t *testing.T) {
	employeeId := 1
	meetingRoom := entity.MeetingRoom{
		Id:          1,
		Name:        "Room A",
		Capacity:    100,
		RatePerDay:  1000000,
		RatePerHour: 200000,
	}
	photoInvoice := "invoice.jpg"
	startDate := time.Now().Add(5 * 24 * time.Hour)
	endDate := startDate.Add(3 * time.Hour)
	contactPIC := "081313131313"
	booking, _ := entity.NewBooking(employeeId, meetingRoom, photoInvoice, contactPIC, startDate, endDate)
	err := booking.Cancel()
	require.Error(t, err)
	assert.Equal(t, "BOOK", booking.GetStatus())
}

func Test_ItShouldBeSuccessAddDiscount(t *testing.T) {
	employeeId := 1
	meetingRoom := entity.MeetingRoom{
		Id:          1,
		Name:        "Room A",
		Capacity:    100,
		RatePerDay:  1000000,
		RatePerHour: 200000,
	}
	photoInvoice := "invoice.jpg"
	startDate := time.Now().Add(8 * 24 * time.Hour)
	endDate := startDate.Add(3 * time.Hour)
	contactPIC := "081313131313"
	booking, _ := entity.NewBooking(employeeId, meetingRoom, photoInvoice, contactPIC, startDate, endDate)
	discount := 50000
	err := booking.AddDiscount(discount)
	require.NoError(t, err)
	assert.Equal(t, discount, booking.GetDiscount())
	assert.Equal(t, 550000, booking.GetGrandTotal())
}

func Test_ItShouldBeFailedAddDiscountWhenDiscountMoreThanGrandTotal(t *testing.T) {
	employeeId := 1
	meetingRoom := entity.MeetingRoom{
		Id:          1,
		Name:        "Room A",
		Capacity:    100,
		RatePerDay:  1000000,
		RatePerHour: 200000,
	}
	photoInvoice := "invoice.jpg"
	startDate := time.Now().Add(8 * 24 * time.Hour)
	endDate := startDate.Add(3 * time.Hour)
	contactPIC := "081313131313"
	booking, _ := entity.NewBooking(employeeId, meetingRoom, photoInvoice, contactPIC, startDate, endDate)
	discount := 5000000
	err := booking.AddDiscount(discount)
	require.Error(t, err)
}

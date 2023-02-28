package booking_http_delivery_test

import (
	"bytes"
	"encoding/json"
	"github.com/agisnur24/book-meetingroom-management-system.git/domain/entity"
	delivery_http "github.com/agisnur24/book-meetingroom-management-system.git/internal/delivery/http"
	booking_http_delivery "github.com/agisnur24/book-meetingroom-management-system.git/internal/delivery/http/booking"
	"github.com/agisnur24/book-meetingroom-management-system.git/internal/delivery/http/request"
	"github.com/agisnur24/book-meetingroom-management-system.git/internal/delivery/http/response"
	"github.com/agisnur24/book-meetingroom-management-system.git/internal/usecase/booking"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestItShouldBeReturnSuccessNewBooking(t *testing.T) {

	meetingRoomId := 1
	employeeId := 1
	picContactInformation := "08121211221"
	requestBody := request.NewBookingRequest{
		MeetingRoomId:         meetingRoomId,
		PicContactInformation: picContactInformation,
		StartDateTime:         "2022-08-30 09:00:00",
		EndDateTime:           "2022-08-30 11:00:00",
	}
	startDatetimeParse, _ := time.Parse("2006-01-02 15:04:05", requestBody.StartDateTime)
	startDatetime := startDatetimeParse
	endDatetimeParse, _ := time.Parse("2006-01-02 15:04:05", requestBody.EndDateTime)
	endDatetime := endDatetimeParse
	usecase := new(booking.UseCaseBookingMock)
	meetingRoom := entity.MeetingRoom{
		Id:          1,
		Name:        "Cendrawasih",
		Capacity:    100,
		RatePerDay:  1000000,
		RatePerHour: 20000,
	}
	expectBooking, _ := entity.NewBooking(employeeId, meetingRoom, "invoice.jpg", picContactInformation,
		startDatetime, endDatetime)
	usecase.On("BookMeetingRoom",
		meetingRoomId, employeeId, startDatetime, endDatetime, picContactInformation).
		Return(expectBooking, nil)

	handler := booking_http_delivery.NewBookingHttpDelivery(usecase)

	requestBodyJson, _ := json.Marshal(requestBody)
	req, err := http.NewRequest("POST", "/booking", bytes.NewBuffer(requestBodyJson))
	require.NoError(t, err)
	recorder := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/booking", handler.NewBooking).Methods("POST")
	router.ServeHTTP(recorder, req)

	require.Equal(t, http.StatusCreated, recorder.Code)
	expectedBookingResponse := response.BookingResponse{
		MeetingRoomId:         meetingRoomId,
		PicContactInformation: picContactInformation,
		StartDateTime:         startDatetime.Format("2006-01-02 15:04:05"),
		EndDateTime:           endDatetime.Format("2006-01-02 15:04:05"),
		Total:                 expectBooking.GetTotal(),
		Discount:              expectBooking.GetDiscount(),
		GrandTotal:            expectBooking.GetGrandTotal(),
	}
	expectedjsonResponse, _ := json.Marshal(expectedBookingResponse)
	assert.Equal(t, string(expectedjsonResponse), recorder.Body.String())
}

func Test_ItShouldBeErrorWhenNotValidFormatDate(t *testing.T) {
	meetingRoomId := 1
	picContactInformation := "0812121212"
	requestBody := request.NewBookingRequest{
		MeetingRoomId:         meetingRoomId,
		PicContactInformation: picContactInformation,
		StartDateTime:         "2022-08-30",
		EndDateTime:           "2022-08-30 11:00:00",
	}
	usecase := new(booking.UseCaseBookingMock)
	handler := booking_http_delivery.NewBookingHttpDelivery(usecase)
	requestBodyJson, _ := json.Marshal(requestBody)
	req, err := http.NewRequest("POST", "/booking", bytes.NewBuffer(requestBodyJson))
	require.NoError(t, err)
	recorder := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/booking", handler.NewBooking).Methods("POST")
	router.ServeHTTP(recorder, req)

	require.Equal(t, http.StatusBadRequest, recorder.Code)
	expectErr := delivery_http.ErrorResponse{Message: "Invalid format date"}
	expectedjsonResponse, _ := json.Marshal(expectErr)
	assert.Equal(t, string(expectedjsonResponse), recorder.Body.String())
}

package booking_http_delivery

import (
	"encoding/json"
	"github.com/agisnur24/book-meetingroom-management-system.git/domain/usecase"
	delivery_http "github.com/agisnur24/book-meetingroom-management-system.git/internal/delivery/http"
	"github.com/agisnur24/book-meetingroom-management-system.git/internal/delivery/http/request"
	"github.com/agisnur24/book-meetingroom-management-system.git/internal/delivery/http/response"
	"net/http"
	"time"
)

type BookingHttpDelivery struct {
	useCase usecase.BookingUseCase
}

func NewBookingHttpDelivery(useCase usecase.BookingUseCase) BookingHttpDelivery {
	return BookingHttpDelivery{useCase: useCase}
}

func (d BookingHttpDelivery) NewBooking(w http.ResponseWriter, r *http.Request) {
	var req request.NewBookingRequest
	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(&req)

	startDatetime, err := time.Parse("2006-01-02 15:04:05", req.StartDateTime)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		errResponse := delivery_http.ErrorResponse{Message: "Invalid format date"}
		errJsonResponse, _ := json.Marshal(errResponse)
		w.Write(errJsonResponse)
		return
	}
	endDatetime, _ := time.Parse("2006-01-02 15:04:05", req.EndDateTime)
	booking, _ := d.useCase.BookMeetingRoom(req.MeetingRoomId, 1,
		startDatetime, endDatetime, req.PicContactInformation)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	res := response.BookingResponse{
		MeetingRoomId:         booking.MeetingRoom.Id,
		PicContactInformation: booking.ContactPIC,
		StartDateTime:         booking.StartDatetime.Format("2006-01-02 15:04:05"),
		EndDateTime:           booking.EndDatetime.Format("2006-01-02 15:04:05"),
		Total:                 booking.GetTotal(),
		Discount:              booking.GetDiscount(),
		GrandTotal:            booking.GetGrandTotal(),
	}

	json, _ := json.Marshal(res)
	w.Write(json)
}

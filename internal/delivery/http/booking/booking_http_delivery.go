package booking_http_delivery

import (
	"book-meeting-hotel/domain/usecase"
	http2 "book-meeting-hotel/internal/delivery/http"
	"encoding/json"
	"net/http"
	"time"
)

type BookingHttpDelivery struct {
	usecase usecase.BookingUseCase
}

func NewBookingHttpDelivery(usecase usecase.BookingUseCase) BookingHttpDelivery {
	return BookingHttpDelivery{usecase: usecase}
}

func (d BookingHttpDelivery) NewBooking(w http.ResponseWriter, r *http.Request) {
	var req NewBookingRequest
	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(&req)

	startDatetime, err := time.Parse("2006-01-02 15:04:05", req.StartDateTime)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		errResponse := http2.ErrorResponse{Message: "Invalid format date"}
		errJsonResponse, _ := json.Marshal(errResponse)
		w.Write(errJsonResponse)
		return
	}
	endDatetime, _ := time.Parse("2006-01-02 15:04:05", req.EndDateTime)
	booking, _ := d.usecase.BookMeetingRoom(req.MeetingRoomId, 1,
		startDatetime, endDatetime, req.PicContactInformation)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := BookingResponse{
		MeetingRoomId:         booking.MeetingRoom.Id,
		PicContactInformation: booking.ContactPIC,
		StartDateTime:         booking.StartDatetime.Format("2006-01-02 15:04:05"),
		EndDateTime:           booking.EndDatetime.Format("2006-01-02 15:04:05"),
		Total:                 booking.GetTotal(),
		Discount:              booking.GetDiscount(),
		GrandTotal:            booking.GetGrandTotal(),
	}

	json, _ := json.Marshal(response)
	w.Write(json)
}

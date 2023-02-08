package booking_http_delivery

type NewBookingRequest struct {
	MeetingRoomId         int    `json:"'meeting_room_id'"`
	PicContactInformation string `json:"pic_contact_information"`
	StartDateTime         string `json:"start_date_time"`
	EndDateTime           string `json:"end_date_time"`
}

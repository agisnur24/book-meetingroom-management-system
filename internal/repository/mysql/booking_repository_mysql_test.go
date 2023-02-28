package mysql_test

import (
	"bou.ke/monkey"
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/agisnur24/book-meetingroom-management-system.git/domain/entity"
	"github.com/agisnur24/book-meetingroom-management-system.git/internal/repository/mysql"
	"github.com/agisnur24/book-meetingroom-management-system.git/internal/repository/mysql/mapper"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type BookingRepositoryMysqlTestSuite struct {
	mock sqlmock.Sqlmock
	db   *sql.DB
	repo mysql.BookingRepositoryMysql
	suite.Suite
}

func (s *BookingRepositoryMysqlTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	require.NoError(s.T(), err)
	s.mock = mock
	s.db = db
	s.repo = mysql.NewBookingRepositoryMysql(db, mapper.BookingMapper{})
}

func (s *BookingRepositoryMysqlTestSuite) TearDownTest() {
	s.db.Close()
}

func Test_BookingRepositoryMysqlSuite(t *testing.T) {
	suite.Run(t, new(BookingRepositoryMysqlTestSuite))
}

func (s *BookingRepositoryMysqlTestSuite) Test_ItShouldBeSuccessSaveBooking() {
	employeeId := 1
	meetingRoom := entity.MeetingRoom{
		Id:          1,
		Name:        "Test",
		Capacity:    11,
		RatePerDay:  30,
		RatePerHour: 30,
	}
	startDate := time.Now().AddDate(0, 0, 10)
	endDate := startDate.AddDate(0, 0, 10)
	booking, _ := entity.NewBooking(employeeId, meetingRoom, "invoice.jpg", "0812121212",
		startDate, endDate)
	s.mock.ExpectExec("INSERT INTO booking").WithArgs(booking.Id, booking.EmployeeId,
		booking.MeetingRoom.Id, booking.PhotoInvoice, booking.ContactPIC, booking.StartDatetime, booking.EndDatetime,
		booking.Status, booking.RatePerDay, booking.RatePerHour, booking.GetDiscount(), booking.BookDate).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := s.repo.Save(booking)
	require.NoError(s.T(), err)
	require.NoError(s.T(), s.mock.ExpectationsWereMet())

}

func (s *BookingRepositoryMysqlTestSuite) Test_ItShouldBeErrorWhenDBDown() {
	employeeId := 1
	meetingRoom := entity.MeetingRoom{
		Id:          1,
		Name:        "Test",
		Capacity:    11,
		RatePerDay:  30,
		RatePerHour: 30,
	}
	startDate := time.Now().AddDate(0, 0, 10)
	endDate := startDate.AddDate(0, 0, 10)
	booking, _ := entity.NewBooking(employeeId, meetingRoom, "invoice.jpg", "0812121212",
		startDate, endDate)
	s.mock.ExpectExec("INSERT INTO booking").WithArgs(booking.Id, booking.EmployeeId,
		booking.MeetingRoom.Id, booking.PhotoInvoice, booking.ContactPIC, booking.StartDatetime, booking.EndDatetime,
		booking.Status, booking.RatePerDay, booking.RatePerHour, booking.GetDiscount(), booking.BookDate).
		WillReturnError(errors.New("mysql down"))

	err := s.repo.Save(booking)
	require.Error(s.T(), err)
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func (s *BookingRepositoryMysqlTestSuite) Test_ItShouldBeSuccessRetrieveBookingByDate() {
	employeeId := 1
	meetingRoom := entity.MeetingRoom{
		Id:          1,
		Name:        "Test",
		Capacity:    11,
		RatePerDay:  30,
		RatePerHour: 30,
	}
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2022, 8, 25, 20, 49, 0, 0, time.Local)
	})
	startDate := time.Now().AddDate(0, 0, 10)
	endDate := startDate.AddDate(0, 0, 10)
	booking, _ := entity.NewBooking(employeeId, entity.MeetingRoom{}, "invoice.jpg", "0812121212",
		startDate, endDate)
	booking.Id = 1
	rows := s.mock.NewRows([]string{"id", "employee_id", "meeting_room_id", "photo_invoice", "contact_pic",
		"start_datetime", "end_datetime", "status", "rate_per_day", "rate_per_hour", "book_date"}).
		AddRow(booking.Id, employeeId, meetingRoom.Id, booking.PhotoInvoice, booking.ContactPIC,
			booking.StartDatetime,
			booking.EndDatetime, booking.Status, booking.RatePerDay, booking.RatePerHour,
			booking.BookDate)
	s.mock.ExpectQuery("SELECT (.+) FROM booking WHERE start_datetime > (.+) and end_datetime < (.+)").
		WithArgs(startDate, endDate).
		WillReturnRows(rows)

	retrieveBooking, err := s.repo.GetByDateAndMeetingRoom(meetingRoom.Id, startDate, endDate)
	require.NoError(s.T(), err)
	require.Equal(s.T(), booking, retrieveBooking)
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

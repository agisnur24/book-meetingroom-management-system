package mysql

import (
	"book-meeting-hotel/domain/entity"
	"book-meeting-hotel/internal/repository/mysql/mapper"
	"book-meeting-hotel/internal/repository/mysql/model"
	"context"
	"database/sql"
	"fmt"
	"github.com/rocketlaunchr/dbq/v2"
	"time"
)

type BookingRepositoryMysql struct {
	db     *sql.DB
	mapper mapper.BookingMapper
}

func NewBookingRepositoryMysql(db *sql.DB, mapper mapper.BookingMapper) BookingRepositoryMysql {
	return BookingRepositoryMysql{db: db, mapper: mapper}
}

func (r BookingRepositoryMysql) Save(booking *entity.Booking) error {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	insertData := []interface{}{
		dbq.Struct(r.mapper.FromDomain(booking)),
	}

	stmt := dbq.INSERTStmt(model.BookingMysqlModel{}.GetTableName(), model.FieldTable(), len(insertData))
	_, err := dbq.E(ctx, r.db, stmt, nil, insertData)
	return err
}

func (r BookingRepositoryMysql) GetByDateAndMeetingRoom(meetingRoomId int, startDatetime time.Time,
	endDatetime time.Time) (*entity.Booking, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	opts := &dbq.Options{ConcreteStruct: model.BookingMysqlModel{}, SingleResult: true,
		DecoderConfig: dbq.StdTimeConversionConfig()}
	query := fmt.Sprintf("SELECT * FROM %s WHERE start_datetime > ? and end_datetime < ?",
		model.BookingMysqlModel{}.GetTableName())
	result, err := dbq.Q(ctx, r.db, query, opts, startDatetime, endDatetime)

	if err != nil {
		return nil, err
	}
	return r.mapper.ToDomain(result.(*model.BookingMysqlModel)), nil
}

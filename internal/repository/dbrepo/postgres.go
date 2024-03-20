package dbrepo

import (
	"context"
	"fmt"
	"time"

	"github.com/thiruthanikaiarasu/udemy-go/bookings/internal/models"
)

func (m *postgresDBRepo) AllUsers() bool {
	return true
}

// InsertReservation inserts a reservation into the database
func (m *postgresDBRepo) InsertReservation(res models.Reservation) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var newId int

	stmt := `insert into reservations (first_name, last_name, email, phone, start_date, end_date, room_id, created_at, updated_at)
			values ($1, $2, $3, $4, $5, $6, $7, $8, $9) returning id
	`

	err := m.DB.QueryRowContext(ctx, stmt, // ExecContext only execute (statement), with QueryRowContext (query) we can scan the value from the query
		res.FirstName,
		res.LastName,
		res.Email,
		res.Phone,
		res.StartDate,
		res.EndDate,
		res.RoomID,
		time.Now(),
		time.Now(),
	).Scan(&newId)
	if err != nil {
		return 0, err
	}

	return newId, nil
}

// InsertRoomRestriction inserts a room restriction into the database
func (m *postgresDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `insert into rooms_restrictions (start_date, end_date, room_id, reservation_id, created_at, updated_at, restriction_id)
			values ($1, $2, $3, $4, $5, $6, $7)`

	_, err := m.DB.ExecContext(ctx, stmt,
		r.StartDate,
		r.EndDate,
		r.RoomID,
		r.ReservationID,
		time.Now(),
		time.Now(),
		r.RestrictionID,
	)
	if err != nil {
		return err
	}

	return nil
}

// SearchAvailabilityByDatesByRoomID returns true if availability exists of roomId, and false if no availability exist
func (m *postgresDBRepo) SearchAvailabilityByDatesByRoomID(start, end time.Time, roomId int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var numRows int

	query := `
		SELECT
			COUNT(id)
		FROM 
			rooms_restrictions
		WHERE 
			room_id = $1 AND 
			$2 BETWEEN start_date AND end_date AND 
			$3 BETWEEN start_date AND end_date;
`
	
	row := m.DB.QueryRowContext(ctx, query, roomId, start, end)
	err := row.Scan(&numRows)
	fmt.Printf("Start date : %T end date : %T ", start, end)
	fmt.Printf("Rooms available : %d", numRows)
	if err != nil {
		return false, err
	}

	return numRows == 0, nil 
}

// SearchAvailabilityForAllRooms returns a slice of available rooms, if any, for given date ranges 
func (m *postgresDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var rooms []models.Room

	query := `
		SELECT 
			r.id, r.room_name
		FROM 
			rooms r
		WHERE r.id not in 
		(SELECT 
			room_id 
		FROM 
			rooms_restrictions rr 
		WHERE 
			$1 BETWEEN rr.start_date AND rr.end_date AND 
			$2 BETWEEN rr.start_date AND rr.end_date AND 
			rr.room_id = r.id)

`
	rows, err := m.DB.QueryContext(ctx, query, start, end)
	if err != nil {
		return rooms, err
	}

	for rows.Next() {
		var room models.Room
		err := rows.Scan(
			&room.ID,
			&room.RoomName,
		)
		if err != nil {
			return rooms, err
		}
		fmt.Printf("Rooms available : %T", rooms)
		rooms = append(rooms, room)
	}

	if err = rows.Err(); err != nil{
		return rooms, nil 
	} 

	return rooms, nil 

}

// GetRoomById gets the room by id 
func (m *postgresDBRepo) GetRoomById(id int) (models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var room models.Room

	query := ` 
		select id, room_name, created_at, updated_at from rooms where id = $1	
`
	row := m.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&room.ID,
		&room.RoomName,
		&room.CreatedAt,
		&room.UpdatedAt,
	)

	if err != nil {
		return room, err
	}

	return room, nil 
}
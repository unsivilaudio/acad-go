package models

import (
	"demoapi/db"
	"time"
)

type Event struct {
	ID          int       `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"date_time" binding:"required"`
	UserID      int64     `json:"user_id"`
}

func (e *Event) Save() error {
	query := `
		INSERT INTO events (
			name,
			description,
			location,
			date_time,
			user_id
		)
		VALUES (?,?,?,?,?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(&e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.ID = int(id)
	return err
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return []Event{}, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		rows.Scan(
			&event.ID,
			&event.Name,
			&event.Location,
			&event.Description,
			&event.DateTime,
			&event.UserID,
		)

		events = append(events, event)
	}

	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(
		&event.ID,
		&event.Name,
		&event.Location,
		&event.Description,
		&event.DateTime,
		&event.UserID,
	)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (e *Event) Update() error {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, dateTime = ?
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		&e.Name,
		&e.Description,
		&e.Location,
		&e.DateTime,
		&e.ID,
	)
	return err
}

func (e *Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(e.ID)
	return err
}

func (e *Event) Register(userId int64) error {
	query := `
		INSERT INTO registrations(event_id, user_id)
		VALUES (?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)
	return err
}

func (e *Event) CancelRegistration(userID int64) error {
	query := `
		DELETE FROM registrations
		WHERE event_id = ? AND user_id = ?
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userID)
	return err
}

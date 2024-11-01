package models

import (
	"time"

	"example.com/webserver/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

// var events = []Event{}

func (e *Event) Save() error {
	query := `INSERT INTO events(name, description,  location, datetime, user_id)
	VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	// place holder '?'
	// kalo mau update database pake exec ama prepare
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id
	return err
	// events = append(events, e)
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	// kalo cm fetch pake query
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		// HARUS SESUAI URUTAN DENGAN STRUCT EVENT
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := `SELECT * FROM events WHERE id = ?`
	// queryrow karena cuma mau fetch 1 row
	row := db.DB.QueryRow(query, id)

	var event Event

	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if err != nil {
		return nil, err
	}

	// if err == sql.ErrNoRows {
	//     return Event{}, fmt.Errorf("event not found with id %d", id)
	// } else if err!= nil {
	//     return Event{}, err
	// }

	return &event, nil
}

func (event Event) Update() error {
	query := `
	UPDATE events 
	SET name =?, description =?, location =?, datetime =?
	WHERE id =?
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	// exec buat update database
	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	if err != nil {
		return err
	}

	return nil
}

func (event Event) Delete() error {
	query := `DELETE FROM events WHERE id =?`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	// exec buat update(delete) data
	_, err = stmt.Exec(event.ID)
	if err != nil {
		return err
	}

	return nil
}

func (e Event) Register(userId int64) error {
	query := "INSERT INTO registrations(event_id, user_id) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)

	return err
}

func (e Event) CancelRegistration(userId int64) error {
	query := "DELETE FROM registrations WHERE event_id =? AND user_id =?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)

	return err
}

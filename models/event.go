package models

import (
	"log"
	"modern-delivery-service/db"
	"time"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

var events []Event = []Event{}

func (event *Event) Save() error {
	log.Printf("Saving event: %+v\n", event)
	query := `
INSERT INTO events (name, description, location, dateTime, user_id)
values ($1, $2, $3, $4, $5)
RETURNING id
`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		log.Printf("Error Prepare: %v\n", err)
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(event.Name, event.Description, event.Location, event.DateTime, event.UserID).Scan(&event.ID)
	if err != nil {
		log.Printf("Error Exec: %v\n", err)
		return err
	}
	return err
}

func GetAllEvents() ([]Event, error) {
	query := `
SELECT id, name, description, location, dateTime, user_id FROM events
`
	rows, err := db.DB.Query(query)
	if err != nil {
		log.Printf("Error querying database: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			log.Printf("Error scanning row: %v\n", err)
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = $1"
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (event Event) Update() error {
	query := `
UPDATE events 
SET name = $2, description = $3, location = $4, dateTime = $5
WHERE id = $1
`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		log.Printf("Error Prepare: %v\n", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(event.ID, event.Name, event.Description, event.Location, event.DateTime)
	if err != nil {
		log.Printf("Error Exec: %v\n", err)
		return err
	}
	return err
}

func (event Event) Delete() error {
	query := "DELETE FROM events WHERE id = $1"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(event.ID)
	if err != nil {
		log.Printf("Error Exec: %v\n", err)
		return err
	}
	return err
}

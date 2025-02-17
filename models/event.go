package models

import (
	"fmt"
	"time"

	"github.com/singhJasvinder101/db"
)

type Event struct {
	Id int64
	Title string		`binding:"required"`
	Description string	`binding:"required"`
	Location string		`binding:"required"`
	DateTime time.Time
	UserId int64	//creator of event
}


func (e *Event) Save() error {
	query := `INSERT INTO events 
	(title, description, location, date_time, user_id) 
	VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)	// prepare for later executions (concurrent)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(e.Title, e.Description, e.Location, e.DateTime, e.UserId)

	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	e.Id = id
	return err
	// events = append(events, e)
}

func (e Event) Update () error {
	query := `UPDATE events SET title = ?, description = ?, location = ?, date_time = ? WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.Title, e.Description, e.Location, e.DateTime, e.Id)
	return err
}

func GetallEvents() ([]Event, error) {
	fmt.Println("Fetching events...")
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	events := []Event{}
	for rows.Next() {
		var e Event
		var dateTimeStr string 

		// dateTime in db stored as string
		err := rows.Scan(&e.Id, &e.Title, &e.Description, &e.Location, &dateTimeStr, &e.UserId)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}

		// parse dateTimeStr to time.Time for retriveing by go
		parsedTime, err := time.Parse("2006-01-02 15:04:05-07:00", dateTimeStr)
		if err != nil {
			fmt.Println("Error parsing date_time:", err)
			return nil, err
		}
		e.DateTime = parsedTime

		fmt.Println("Fetched event:", e)
		events = append(events, e)
	}

	return events, nil
}

func DeleteAllRows () error {
	if db.DB == nil {
        return fmt.Errorf("database connection is nil")
    }
	query := `DELETE FROM events`
	_, err := db.DB.Exec(query)
	return err
}

func GetEventById(id string) (Event, error){
	query := `select * from events where id = ?`
	row := db.DB.QueryRow(query, id)
	var e Event
	var dateTimeStr string

	row.Scan(&e.Id, &e.Title, &e.Description, &e.Location, &dateTimeStr, &e.UserId)
	parsedTime, err := time.Parse("2006-01-02 15:04:05-07:00", dateTimeStr)
	if err != nil {
		fmt.Println("Error parsing date_time:", err)
		return Event{}, err
	}
	e.DateTime = parsedTime
	return e, nil
}

func DeleteEventById(id string) error {
	query := `DELETE FROM events WHERE id = ?`
	_, err := db.DB.Exec(query, id)
	return err
}



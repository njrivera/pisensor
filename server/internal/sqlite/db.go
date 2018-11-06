package sqlite

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pisensor/pkg/models"
)

const (
	insertReadingStmt       = "INSERT INTO readings (serial, model, temp, unit, timestamp) VALUES (?,?,?,?,?)"
	getReadingsBetweenTimes = "SELECT model, temp, unit, timestamp FROM readings WHERE serial = ? AND timestamp BETWEEN ? AND ?"

	sqliteTimeFormat = "2006-01-02 15:04:05-07:00"
)

type SQLiteDB struct {
	conn *sql.DB
}

func NewDB() (*SQLiteDB, error) {
	conn, err := sql.Open("sqlite3", GetDBPath())
	if err != nil {
		return nil, err
	}

	return &SQLiteDB{
		conn: conn,
	}, nil
}

func (db *SQLiteDB) GetTempsBetweenTimes(serial string, start, end time.Time) ([]models.TempReading, error) {
	rows, err := db.conn.Query(getReadingsBetweenTimes, serial, start, end)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	readings := []models.TempReading{}

	for rows.Next() {
		var r models.TempReading
		var timeStr string
		if err := rows.Scan(&r.Model, &r.Temp, &r.Unit, &timeStr); err != nil {
			return nil, err
		}

		if r.Timestamp, err = time.Parse(sqliteTimeFormat, timeStr); err != nil {
			return nil, err
		}

		readings = append(readings, r)
	}

	return readings, nil
}

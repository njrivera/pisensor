package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pisensor/pkg/models"
)

const (
	insertReadingStmt = "INSERT INTO readings (serial, model, temp, unit, timestamp) VALUES (?,?,?,?,?)"
)

type SQLiteDB struct {
	conn *sql.DB
	stmt *sql.Stmt
}

func NewDB() (*SQLiteDB, error) {
	conn, err := sql.Open("sqlite3", GetDBPath())
	if err != nil {
		return nil, err
	}

	stmt, err := conn.Prepare(insertReadingStmt)
	if err != nil {
		return nil, err
	}

	return &SQLiteDB{
		conn: conn,
		stmt: stmt,
	}, nil
}

func (db *SQLiteDB) InsertReading(reading models.TempReading) error {
	_, err := db.stmt.Exec(
		reading.Serial,
		reading.Model,
		reading.Temp,
		reading.Unit,
		reading.Timestamp,
	)

	return err
}

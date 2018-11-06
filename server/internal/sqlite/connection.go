package sqlite

import (
	"os"
	"sync"
)

const (
	dbPathEnv = "TEMP_READINGS_SQLITE_PATH"
)

var (
	dbOnce sync.Once
	dbPath string
)

func GetDBPath() string {
	dbOnce.Do(func() {
		dbPath = os.Getenv(dbPathEnv)
		if dbPath == "" {
			dbPath = "/opt/pisensor/temp_reading.db"
		}
	})

	return dbPath
}

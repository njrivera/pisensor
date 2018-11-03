package dbfactory

import (
	"fmt"

	"github.com/pisensor/recorder/internal/sqlite"
	"github.com/pisensor/recorder/internal/types"
)

const (
	SQLite = "SQLITE"
)

type ErrUnknownDBType struct {
	dbType string
}

func (e ErrUnknownDBType) Error() string {
	return fmt.Sprintf("Unknown DB type: %s", e.dbType)
}

func NewDB(dbType string) (types.DB, error) {
	switch dbType {
	case SQLite:
		return sqlite.NewDB()
	default:
		return nil, ErrUnknownDBType{dbType: dbType}
	}
}

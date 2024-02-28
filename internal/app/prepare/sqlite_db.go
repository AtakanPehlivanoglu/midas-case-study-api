package prepare

import (
	"database/sql"
)

func SqliteDb(driverName, dbName string) (*sql.DB, error) {
	db, err := sql.Open(driverName, dbName)
	if err != nil {
		return nil, err
	}
	return db, nil
}

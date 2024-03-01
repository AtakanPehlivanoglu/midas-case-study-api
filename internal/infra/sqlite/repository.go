package sqliterepo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/AtakanPehlivanoglu/midas-case-study-api/internal/domain"
	_ "github.com/mattn/go-sqlite3"
)

type SqliteDb struct {
	db *sql.DB
}

func NewSqliteRepository(args SqliteRepositoryArgs) (*SqliteDb, error) {
	if args.Db == nil {
		return nil, errors.New("db is missing")
	}
	return &SqliteDb{
		db: args.Db,
	}, nil
}

type SqliteRepositoryArgs struct {
	Db *sql.DB
}

// CreateFile creates file entity for the table.
func (s *SqliteDb) CreateFile(ctx context.Context, args domain.CreateFileArgs) error {
	db := s.db

	sqlStmt := `
	create table if not exists file (name text not null, content text);
	`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		return fmt.Errorf("%q: %s\n", err, sqlStmt)
	}

	_, err = db.Exec(args.Query)
	if err != nil {
		return fmt.Errorf("%q: %s\n", err, args.Query)
	}

	return nil
}

// CreateFileQueryBuilder creates initial insert statement for bulk insert.
func CreateFileQueryBuilder() string {
	return fmt.Sprintf("insert into file(name, content) values")
}

// CreateFileBulkInsertBuilder creates additional insert row for bulk insert.
func CreateFileBulkInsertBuilder(fileName, fileContent string) string {
	// build multi line insert
	return fmt.Sprintf("('%v', '%v'),", fileName, fileContent)
}

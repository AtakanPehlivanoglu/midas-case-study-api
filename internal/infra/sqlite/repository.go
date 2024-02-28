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

func (s *SqliteDb) CreateFile(ctx context.Context, args domain.CreateFileArgs) error {
	db := s.db

	sqlStmt := `
	create table if not exists file (name text not null, content text);
	`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		return fmt.Errorf("%q: %s\n", err, sqlStmt)
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("insert into file(name, content) values(?, ?)")
	if err != nil {
		return fmt.Errorf("%q: %s\n", err, sqlStmt)
	}
	defer stmt.Close()
	for i := 10; i < 20; i++ {
		_, err = stmt.Exec(i, fmt.Sprintf("test data here: %v", i))
		if err != nil {
			return fmt.Errorf("%q: %s\n", err, sqlStmt)
		}
	}
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("%q: %s\n", err, sqlStmt)
	}

	return nil
}

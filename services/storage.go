package services

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"

	"client/interfaces"
)

type Storage struct {
	db     *sql.DB
	dbPath *string
	dbName *string
	tables *string
}

func (s *Storage) New(dbPath *string, dbName *string, tables *string) bool {
	s.dbPath = dbPath
	s.dbName = dbName
	s.tables = tables
	return true
}

func (s *Storage) Initialize() {
	var err error
	defer s.db.Close()
	_, err = s.db.Exec(*s.tables)
	if err != nil {
		panic(err)
	}
}

func (s *Storage) Open() bool {

	folderPath := *s.dbPath
	dbPath := filepath.Join(folderPath, *s.dbName)

	if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
		log.Printf("Failed to create directory: %v", err)
		return false
	}

	var err error
	s.db, err = sql.Open("sqlite3", dbPath)
	return err == nil
}

func (s *Storage) Query(sql *string, parameters *[]interface{}) (interfaces.RowsScanner, error) {
	rows, err := s.db.Query(*sql, *parameters...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (s *Storage) QuerySingle(sql *string, parameters *[]interface{}) interfaces.RowScanner {
	row := s.db.QueryRow(*sql, *parameters...)

	return row
}

func (s *Storage) Exec(sql *string, parameters *[]interface{}) error {
	result, err := s.db.Exec(*sql, *parameters...)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()

	if affected == 0 {
		return errors.New("no rows affected")
	}

	return err
}

func (s *Storage) ExecReturnID(sql *string, parameters *[]interface{}) (int, error) {
	result, err := s.db.Exec(*sql, *parameters...)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	affected, err := result.RowsAffected()

	if affected == 0 {
		return 0, errors.New("no rows affected")
	}

	return int(lastInsertID), err
}

func (s *Storage) Close() bool {
	if s.db != nil {
		s.db.Close()
	}
	return true
}
